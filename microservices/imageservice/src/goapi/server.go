package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/codegangsta/negroni"
	"github.com/satori/go.uuid"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initConfig()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// Initialize Config Variables
func initConfig() {
	initSecretFromEnv()
	initConfigFileFromS3()
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/image/{id}", imageHandler(formatter)).Methods("GET")
	mx.HandleFunc("/image", imagePostHandler(formatter)).Methods("POST", "OPTIONS")
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}


func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// API Get Image Handler
func imageHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		enableCors(&w)
		params := mux.Vars(req)
		var imageId string = params["id"]
		fmt.Println( "Image ID: ", imageId )
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
			var errorResponse ErrorResponse
			errorResponse.Message = "Server Error"
			formatter.JSON(w, http.StatusInternalServerError, errorResponse)
			panic(err)
			return
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        conn := session.DB(mongodb_database).C(mongodb_collection)
		var result bson.M
		query := bson.M{"id" : imageId}
		err = conn.Find(query).One(&result)
        if err != nil {
			log.Print(err)
			var errorResponse ErrorResponse
			errorResponse.Message = "Image not found"
			formatter.JSON(w, http.StatusBadRequest, errorResponse)
        } else {
			fmt.Println("Image data:", result )
			var image Image
			bsonBytes, _ := bson.Marshal(result)
			bson.Unmarshal(bsonBytes, &image)
			fmt.Println("Image :", image )
			formatter.JSON(w, http.StatusOK, image)
		}
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// API POST Image Handler
func imagePostHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		setupResponse(&w, req)
		if (*req).Method == "OPTIONS" {
			return
		}
		var image Image
		image.Id = uuid.NewV4().String()
		image.Description = req.FormValue("description")
		image.UserId = req.FormValue("userid")
		image.UserName = req.FormValue("username")

		loc, _ := time.LoadLocation("UTC")
		image.Timestamp = time.Now().In(loc).String()

		foodImageFile, _, err := req.FormFile("foodImage")
		if err != nil {
			fmt.Println("Error in getting the file", err, image,foodImageFile )
			var errorResponse ErrorResponse
			errorResponse.Message = "Invalid Request"
			formatter.JSON(w, http.StatusBadRequest, errorResponse)
			return
		}
		defer foodImageFile.Close()

		// Mongo connection
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			fmt.Println("Error in creating MongoDB session", err)
			var errorResponse ErrorResponse
			errorResponse.Message = "Server Error"
			formatter.JSON(w, http.StatusInternalServerError, errorResponse)
			return
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		// Upload file to S3 Bucket
		resultImageUrl := uploadFileToS3(image.Id, foodImageFile)

		if resultImageUrl == "" {
			fmt.Println("NOT a Valid Request")
			var errorResponse ErrorResponse
			errorResponse.Message = "Server Error"
			formatter.JSON(w, http.StatusInternalServerError, errorResponse)
			return
		} 
		image.Url = resultImageUrl
		err = c.Insert(image)
		if err != nil {
			fmt.Println("Exception inserting data to Database", err)
			var errorResponse ErrorResponse
			errorResponse.Message = "Server Error"
			formatter.JSON(w, http.StatusInternalServerError, errorResponse)
			return
		}
		formatter.JSON(w, http.StatusOK, image)
	}
}
