package main

import (
	"fmt"
	"log"
	"github.com/chasex/redis-go-cluster"
	"time"
	"encoding/json"
)

func saveToTimelineRedis(image_id string, image_details string){
	cluster, err := redis.NewCluster(
		&redis.Options{
		StartNodes: []string{"10.0.1.211:6379", "10.0.1.119:6379", "10.0.1.230:6379"},
		ConnTimeout: 50 * time.Millisecond,
		ReadTimeout: 50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
		KeepAlive: 16,
		AliveTime: 60 * time.Second,
		})

	if err != nil {
		log.Fatalf("redis.New error: %s", err.Error())
		fmt.Println("redis.New error: %s", err.Error())
	}

	cluster.Do("HMSET", "timeline", image_id, image_details)
}

func getTimelineRedis() map[string] string {
	cluster, err := redis.NewCluster(
		&redis.Options{
		StartNodes: []string{"10.0.1.211:6379", "10.0.1.119:6379", "10.0.1.230:6379"},
		ConnTimeout: 50 * time.Millisecond,
		ReadTimeout: 50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
		KeepAlive: 16,
		AliveTime: 60 * time.Second,
		})

	if err != nil {
		log.Fatalf("redis.New error: %s", err.Error())
		fmt.Println("redis.New error: %s", err.Error())
	}

	reply, err2 := redis.StringMap(cluster.Do("HGETALL", "timeline"))

	if err2 != nil {
		log.Fatalf("redis.New error: %s", err.Error())
		fmt.Println("redis.New error: %s", err.Error())
	}

	return reply
}

func updateCommentCountRedis(image_id string, count int64) {
	cluster, err := redis.NewCluster(
		&redis.Options{
		StartNodes: []string{"10.0.1.211:6379", "10.0.1.119:6379", "10.0.1.230:6379"},
		ConnTimeout: 50 * time.Millisecond,
		ReadTimeout: 50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
		KeepAlive: 16,
		AliveTime: 60 * time.Second,
		})

	if err != nil {
		log.Fatalf("redis.New error: %s", err.Error())
		fmt.Println("redis.New error: %s", err.Error())
	}

	reply, err2 := redis.StringMap(cluster.Do("HGETALL", "timeline"))

	if err2 != nil {
		log.Fatalf("redis.New error: %s", err.Error())
		fmt.Println("redis.New error: %s", err.Error())
	}

	image_details := reply[image_id]
	var img image
	bytes := []byte(image_details)
	json.Unmarshal(bytes,&img)
	img.CommentCount = count

	b, _ := json.Marshal(img)
	s := string(b)
	reply[image_id] = s

	cluster.Do("HMSET", "timeline", image_id, reply)
}

func updateLikeCountRedis(image_id string, count int64) {
	cluster, err := redis.NewCluster(
		&redis.Options{
		StartNodes: []string{"10.0.1.211:6379", "10.0.1.119:6379", "10.0.1.230:6379"},
		ConnTimeout: 50 * time.Millisecond,
		ReadTimeout: 50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
		KeepAlive: 16,
		AliveTime: 60 * time.Second,
		})

	if err != nil {
		log.Fatalf("redis.New error: %s", err.Error())
		fmt.Println("redis.New error: %s", err.Error())
	}

	reply, err2 := redis.StringMap(cluster.Do("HGETALL", "timeline"))

	if err2 != nil {
		log.Fatalf("redis.New error: %s", err.Error())
		fmt.Println("redis.New error: %s", err.Error())
	}

	image_details := reply[image_id]
	var img image
	bytes := []byte(image_details)
	json.Unmarshal(bytes,&img)
	img.LikeCount = count

	b, _ := json.Marshal(img)
	s := string(b)
	reply[image_id] = s

	cluster.Do("HMSET", "timeline", image_id, reply)
}

func setKeyRedis(key string, value string) {
	cluster, err := redis.NewCluster(
		&redis.Options{
		StartNodes: []string{"10.0.1.211:6379", "10.0.1.119:6379", "10.0.1.230:6379"},
		ConnTimeout: 50 * time.Millisecond,
		ReadTimeout: 50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
		KeepAlive: 16,
		AliveTime: 60 * time.Second,
		})

	if err != nil {
		log.Fatalf("redis.New error: %s", err.Error())
		fmt.Println("redis.New error: %s", err.Error())
	}
	
	cluster.Do("SET", key, value)
}

func getValueRedis(key string) string {
	cluster, err := redis.NewCluster(
		&redis.Options{
		StartNodes: []string{"10.0.1.211:6379", "10.0.1.119:6379", "10.0.1.230:6379"},
		ConnTimeout: 50 * time.Millisecond,
		ReadTimeout: 50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
		KeepAlive: 16,
		AliveTime: 60 * time.Second,
		})

	if err != nil {
		log.Fatalf("redis.New error: %s", err.Error())
		fmt.Println("redis.New error: %s", err.Error())
	}	

	reply, err2 := redis.String(cluster.Do("GET", key))
	
	if err2 != nil {
		log.Fatalf("redis.New error: %s", err.Error())
		fmt.Println("redis.New error: %s", err.Error())
	}

	return reply
}

