package main

type Image struct {
	Id             	string 	
	Description     string    	
	Url 			string	    
	UserId		 	string
	UserName		string
	Timestamp		string	
}

type ErrorResponse struct {
	Message		string
}

type Config struct {
	Mongo Mongo
	Aws AWS
}

type Mongo struct {
	Url string
	Database string
	Collection string
}

type AWS struct {
	S3ContentBucketName string
	S3ContentPath string
	Region string
	SNSTopicArn string
}
