package main

import (
	"fmt"
	"log"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/codegangsta/negroni"
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
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}

// API Get Image Handler
func imageHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
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