package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// MongoDB Config
var mongodb_server = ""
var mongodb_database = ""
var mongodb_collection = ""

// S3 bucket details
var s3_content_bucket_name = ""
var s3_content_path = ""
var aws_region = ""
var sns_topic = ""

// S3 config details
var s3_config_file_name = ""
var s3_config_region = ""
var aws_access_key = ""
var aws_secret_key = ""
var s3_config_bucket_name = ""
var s3_config_path_name = ""

// Initialize config from environment variables
func initSecretFromEnv() {
	s3_config_file_name = os.Getenv("CONFIG_FILE_NAME")
	s3_config_bucket_name = os.Getenv("S3_CONFIG_BUCKET_NAME")
	s3_config_path_name = os.Getenv("S3_CONFIG_PATH_NAME")
	s3_config_region = os.Getenv("S3_CONFIG_REGION")
	aws_access_key = os.Getenv("AWS_ACCESS_KEY")
	aws_secret_key = os.Getenv("AWS_SECRET_KEY")
}

// Initialize config from S3 config file
func initConfigFileFromS3() {
	item := s3_config_file_name
	file, err := os.Create(item)
    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s3_config_region),
		Credentials: credentials.NewStaticCredentials(aws_access_key, aws_secret_key, "")},
	)

	downloader := s3manager.NewDownloader(sess)

    numBytes, err := downloader.Download(file,
        &s3.GetObjectInput{
			Bucket: aws.String(s3_config_bucket_name),
            Key:    aws.String(path.Join(s3_config_path_name, s3_config_file_name)),
        })
    if err != nil {
        fmt.Println(err)
	}
	
	var config Config
 
	filejson, _ := ioutil.ReadAll(file)
	_ = json.Unmarshal(filejson, &config)

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	mongodb_server = config.Mongo.Url
	mongodb_database = config.Mongo.Database
	mongodb_collection = config.Mongo.Collection
	s3_content_bucket_name = config.Aws.S3ContentBucketName
	s3_content_path = config.Aws.S3ContentPath
	aws_region = config.Aws.Region
	sns_topic = config.Aws.SNSTopicArn
}