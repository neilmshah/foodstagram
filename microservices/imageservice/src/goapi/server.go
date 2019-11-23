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
	mx.HandleFunc("/image", allImagesHandler(formatter)).Methods("GET")
	mx.HandleFunc("/image", imagePostHandler(formatter)).Methods("POST", "OPTIONS")
	mx.HandleFunc("/user/{userId}", userImagesHandler(formatter)).Methods("GET")
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

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, apikey")
}

// API Get Image Handler
func imageHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		setupResponse(&w, req)
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

		image.Timestamp = time.Now().UTC().Unix()

		foodImageFile, _, err := req.FormFile("foodImage")
		if err != nil {
			fmt.Println("Error in getting the file", err, image, foodImageFile)
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
		go publishSNS(image)
		formatter.JSON(w, http.StatusOK, image)
	}
}

// API Get Images For User Handler
func userImagesHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		enableCors(&w)
		params := mux.Vars(req)
		var userId string = params["userId"]
		fmt.Println( "UserId: ", userId )
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
		result := make([]Image, 10, 10)
		query := bson.M{"userid" : userId}
		err = conn.Find(query).All(&result)
        if err != nil {
			log.Print(err)
			var errorResponse ErrorResponse
			errorResponse.Message = "Images not found for user"
			formatter.JSON(w, http.StatusBadRequest, errorResponse)
        } else {
			fmt.Println("Users images:", result )
			formatter.JSON(w, http.StatusOK, result)
		}
	}
}

// API Get All Images Handler
func allImagesHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		enableCors(&w)
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
		result := make([]Image, 10, 10)
		err = conn.Find(nil).All(&result)
        if err != nil {
			log.Print(err)
			var errorResponse ErrorResponse
			errorResponse.Message = "No image found"
			formatter.JSON(w, http.StatusBadRequest, errorResponse)
        } else {
			formatter.JSON(w, http.StatusOK, result)
		}
	}
}
