package main

import (
	"fmt"
	"time"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

func contains(a string, list []string) int64 {
    fmt.Println(a)
    fmt.Println(list)
    for _, b := range list {
        fmt.Println(b)
        if b == a {
            return 1
        }
    }
    return 0
}

func createComment(photo_id string, user_id string, user_name, comment string){
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
    	new_comment := user_comment{user_id, user_name, timestamp, comment}
    	new_photo := data_struct{photo_id, 0, 1, []string{}, []user_comment{new_comment}}
    	c.Insert(new_photo)
    }else{
		change1 := bson.M{"$push":bson.M{"comments":bson.M{"user_id": user_id, "user_name": user_name, "timestamp": timestamp, "comment": comment}}}
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

func readCommentCount(photo_id string) int64{
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
    return data.Comment_count
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
    like_status := contains(user_id, data.Likes)
    fmt.Println(like_status)
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