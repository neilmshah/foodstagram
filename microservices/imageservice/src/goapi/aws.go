package main
import (
	"fmt"
	"log"
	"mime/multipart"
	"encoding/json"
	"path"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/sns"
)

func uploadFileToS3(filename string, file multipart.File) string {
	res := ""
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(aws_region),
		Credentials: credentials.NewStaticCredentials(aws_access_key, aws_secret_key, "")},
	)
	uploader := s3manager.NewUploader(sess)
	result, s3err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3_content_bucket_name),
		Key: aws.String(path.Join(s3_content_path, filename)),
		Body:   file,
	})
	if s3err != nil {
		log.Fatalf("Unable to upload %q to %q, %v", filename, s3_content_bucket_name, err)
	}

	log.Println("Successfully uploaded image location" + result.Location)
	res = result.Location
	return res
}

func publishSNS(image Image) {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(aws_region),
		Credentials: credentials.NewStaticCredentials(aws_access_key, aws_secret_key, "")},
	)

    if err != nil {
        fmt.Println("Aws SNS error:", err)
        return
    }

	client := sns.New(sess)
	image_marshalled, err := json.Marshal(image)
    input := &sns.PublishInput{
        Message:  aws.String(string(image_marshalled)),
        TopicArn: aws.String(sns_topic),
    }

    result, err := client.Publish(input)
    if err != nil {
        fmt.Println("SNS Publish error:", err)
        return
    }

    fmt.Println(result)
}