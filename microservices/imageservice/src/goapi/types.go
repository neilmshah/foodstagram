package main

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
