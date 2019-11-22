package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

func createLike(photo_id string, user_id string){
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
    	new_photo := data_struct{photo_id, 1, 0, []string{user_id}, []user_comment{}}
    	c.Insert(new_photo)
    }else{
		change := bson.M{"$inc":bson.M{"like_count":1}}
		err = c.Update(query, change)
		if err != nil{
			panic(err)
		}
		change2 := bson.M{"$push":bson.M{"likes":user_id}}
		err = c.Update(query, change2)
		if err != nil{
			panic(err)
		}
	}
    fmt.Println("Like incremented")
}

func readLikes(photo_id string) int64{
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
    return data.Like_count
}