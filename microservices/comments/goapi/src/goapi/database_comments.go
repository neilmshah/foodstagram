package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

func createComment(photo_id string, user_id string, comment string){
	session, err := mgo.Dial(mongodb_server)
    if err != nil {
        panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    c := session.DB(mongodb_database).C(mongodb_collection)
    query := bson.M{"photo_id" : photo_id}
    change := bson.M{"$push":bson.M{"comment_list":comment}}
    err = c.Update(query, change)
    if err != nil{
    	panic(err)
    }

}

func readComments(photo_id string) []string{
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
    return data.Comment_list
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