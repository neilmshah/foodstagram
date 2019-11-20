package main

import (
	"fmt"
	"encoding/json"
	"net/http"
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
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/setkey", setKey(formatter)).Methods("GET")
	mx.HandleFunc("/getkey", getKey(formatter)).Methods("GET")
	mx.HandleFunc("/timeline", getTimeline(formatter)).Methods("GET")
	mx.HandleFunc("/addimage/{image_id}", addImage(formatter)).Methods("POST")
	mx.HandleFunc("/updatelikes/{image_id}", updatelikes(formatter)).Methods("POST")
	mx.HandleFunc("/updatecomments/{image_id}", updatecomments(formatter)).Methods("POST")

}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}

//Set Redis Key
func setKey(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		setKeyRedis("foo","bar")
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Key Set!"})
	}
}

//Get Redis Key
func getKey(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		response := getValueRedis("foo")
		formatter.JSON(w, http.StatusOK, struct{ Test string }{response})
	}
}

//Get Timeline
func getTimeline(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var s map[string] string = getTimelineRedis()

		response_array := make([]image, 0, len(s))
		for _,v := range s {
			var img image
			bytes := []byte(v)
			json.Unmarshal(bytes,&img)
			response_array = append(response_array,img)
		}

		formatter.JSON(w, http.StatusOK, struct{ Timeline []image }{response_array})
	}
}


func addImage(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var img image
		_ = json.NewDecoder(req.Body).Decode(&img)

		params := mux.Vars(req)
		var image_id string = params["image_id"]
		fmt.Println(image_id)

		if image_id == "" {
			formatter.JSON(w, http.StatusOK, struct{ Error string }{"Image ID missing in URL"})
		}

		b, _ := json.Marshal(img)
		s := string(b)

		saveToTimelineRedis(image_id, s)

		formatter.JSON(w, http.StatusOK, struct{ Status string }{"Image Details Added"})
	}
}

func updatecomments(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var commentCount count
		_ = json.NewDecoder(req.Body).Decode(&commentCount)

		params := mux.Vars(req)
		var image_id string = params["image_id"]
		fmt.Println(image_id)

		if image_id == "" {
			formatter.JSON(w, http.StatusOK, struct{ Error string }{"Image ID missing in URL"})
		}

		updateCommentCountRedis(image_id, commentCount.num)

		formatter.JSON(w, http.StatusOK, struct{ Status string }{"Comment count updated"})
	}
}

func updatelikes(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var likeCount count
		_ = json.NewDecoder(req.Body).Decode(&likeCount)

		params := mux.Vars(req)
		var image_id string = params["image_id"]
		fmt.Println(image_id)

		if image_id == "" {
			formatter.JSON(w, http.StatusOK, struct{ Error string }{"Image ID missing in URL"})
		}

		updateLikeCountRedis(image_id, likeCount.num)

		formatter.JSON(w, http.StatusOK, struct{ Status string }{"Like count updated"})
	}
}
