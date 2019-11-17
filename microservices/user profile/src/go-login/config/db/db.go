package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDBCollection() (*mongo.Collection, error) {
	// Set client options
	// credential := options.Credential{
	// 	Username: "admin",
	// 	Password: "admin",
	// }
	// clientOptions := options.Client().ApplyURI("mongodb://3.89.155.155:27017").SetAuth(credential)
	//clientOptions := options.Client().ApplyURI("mongodb://admin:password@18.207.2.160:27017/GoLogin")
	//shard 2
	//clientOptions := options.Client().ApplyURI("mongodb://admin:admin@54.161.239.229:27017/admin")
	//Shard 1
	//clientOptions := options.Client().ApplyURI("mongodb://admin:admin@18.206.233.147:27017/admin")
	//Router
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin@100.27.3.107:27017/admin")
	//working locally
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//try on docker locally
	//clientOptions := options.Client().ApplyURI("mongodb://t800:t800t800@ds041248.mlab.com:41248/imageservice")
	//clientOptions := options.Client().ApplyURI("mongodb://admin:admin@127.0.0.1/GoLogin:27017")
	//clientOptions := options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0-4zfux.mongodb.net/GoLogin?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	//client, err := mongo.Connect(context.TODO(), "mongodb://admin:admin@100.25.215.61:27017/admin")
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	collection := client.Database("GoLogin").Collection("users")
	return collection, nil
}