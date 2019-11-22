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
	// mx.HandleFunc("/setkey", setKey(formatter)).Methods("GET")
	// mx.HandleFunc("/getkey", getKey(formatter)).Methods("GET")
	mx.HandleFunc("/timeline", getTimeline(formatter)).Methods("GET")
	mx.HandleFunc("/addimage", addImage(formatter)).Methods("POST")
	mx.HandleFunc("/updatelikes", updatelikes(formatter)).Methods("POST")
	mx.HandleFunc("/updatecomments", updatecomments(formatter)).Methods("POST")

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

//Get Timeline
func getTimeline(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		enableCors(&w)
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
		var snsReq sns
		_ = json.NewDecoder(req.Body).Decode(&snsReq)

		fmt.Println("Decoded Body: ", snsReq)

		if snsReq.Type == "SubscriptionConfirmation"{
			confirmSubscription(snsReq.SubscribeURL)
			return 
		}

		var img image 
		bytes := []byte(snsReq.Message)
		json.Unmarshal(bytes,&img)

		fmt.Println("Image details: ", img)

		b, _ := json.Marshal(img)
		s := string(b)

		saveToTimelineRedis(img.Id, s)

		formatter.JSON(w, http.StatusOK, struct{ Status string }{"Image Details Added"})
	}
}

func updatecomments(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var snsReq sns
		_ = json.NewDecoder(req.Body).Decode(&snsReq)

		fmt.Println("Decoded Body: ", snsReq)

		if snsReq.Type == "SubscriptionConfirmation"{
			confirmSubscription(snsReq.SubscribeURL)
			return 
		}

		var comments count 
		bytes := []byte(snsReq.Message)
		json.Unmarshal(bytes,&comments)

		fmt.Println("Comment details: ", comments)

		//updateCommentCountRedis(comments.Id, comment.Num)

		formatter.JSON(w, http.StatusOK, struct{ Status string }{"Comment count updated"})
	}
}

func updatelikes(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var snsReq sns
		_ = json.NewDecoder(req.Body).Decode(&snsReq)

		fmt.Println("Decoded Body: ", snsReq)

		if snsReq.Type == "SubscriptionConfirmation"{
			confirmSubscription(snsReq.SubscribeURL)
			return 
		}

		var likes count 
		bytes := []byte(snsReq.Message)
		json.Unmarshal(bytes,&likes)

		fmt.Println("Comment details: ", likes)

		//updateLikeCountRedis(likes.Id, likes.Num)

		formatter.JSON(w, http.StatusOK, struct{ Status string }{"Like count updated"})
	}
}

func confirmSubscription(subcribeURL string) {
    response, err := http.Get(subcribeURL)
    if err != nil {
        fmt.Printf("Unbale to confirm subscriptions")
    } else {
        fmt.Printf("Subscription Confirmed sucessfully. %d", response.StatusCode)
    }
}


// //Set Redis Key
// func setKey(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		setKeyRedis("foo","bar")
// 		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Key Set!"})
// 	}
// }

// //Get Redis Key
// func getKey(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		response := getValueRedis("foo")
// 		formatter.JSON(w, http.StatusOK, struct{ Test string }{response})
// 	}
// }
