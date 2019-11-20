package main
import (
	"log"
	"mime/multipart"
	"path"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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