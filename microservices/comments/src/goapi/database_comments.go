package main

import (
	"fmt"
	"time"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func createComment(photo_id string, user_id string, comment string){
	session, err := mgo.Dial(mongodb_server)
    if err != nil {
        panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    c := session.DB(mongodb_database).C(mongodb_collection)
    query := bson.M{"photo_id" : photo_id}
    var result bson.M
    err = c.Find(query).One(&result)
    // loc, _ := time.LoadLocation("UTC")
    timestamp := time.Now().UTC().Unix()
    if err != nil{
    	new_comment := user_comment{user_id, timestamp, comment}
    	new_photo := data_struct{photo_id, 0, 1, []string{}, []user_comment{new_comment}}
    	c.Insert(new_photo)
    }else{
		change1 := bson.M{"$push":bson.M{"comments":bson.M{"user_id": user_id, "timestamp": timestamp, "comment": comment}}}
		change2 := bson.M{"$inc":bson.M{"comment_count":1}}
		err = c.Update(query, change1)
		if err != nil{
			panic(err)
		}
		err = c.Update(query, change2)
		if err != nil{
			panic(err)
		}
	}
}

func readComments(photo_id string, user_id string) modal_struct{
	session, err := mgo.Dial(mongodb_server)
    if err != nil {
        panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    c := session.DB(mongodb_database).C(mongodb_collection)
    query := bson.M{"photo_id" : photo_id}
    var result bson.M
    err = c.Find(query).One(&result)
    if err != nil{
    	panic(err)
    }
    var data data_struct
    dataBytes, _ := bson.Marshal(result)
    bson.Unmarshal(dataBytes, &data)
    fmt.Println(data)
    like_status := stringInSlice(user_id, data.Likes)

    return modal_struct{like_status, data.Comments}
}

func deleteComment(photo_id string, user_id string, comment string){
	session, err := mgo.Dial(mongodb_server)
    if err != nil {
        panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    c := session.DB(mongodb_database).C(mongodb_collection)
    query := bson.M{"photo_id" : photo_id}
    change := bson.M{"$pull":bson.M{"comment_list":comment}}
    err = c.Update(query, change)
    if err != nil{
    	panic(err)
    }else{
    	fmt.Println("Comment Deleted from Document")
    }
}