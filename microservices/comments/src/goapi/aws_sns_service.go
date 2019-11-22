package main

import (
	"fmt"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func publishSNS(photo_id string, count int64, sns_topic string) {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(AWS_REGION),
		Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY, AWS_SECRET_KEY, "")},
	)

    if err != nil {
        fmt.Println("Aws SNS error:", err)
        return
    }

	client := sns.New(sess)
	data := sns_struct{photo_id, count}
	data_marshalled, err := json.Marshal(data)
    input := &sns.PublishInput{
        Message:  aws.String(string(data_marshalled)),
        TopicArn: aws.String(sns_topic),
    }

    result, err := client.Publish(input)
    if err != nil {
        fmt.Println("SNS Publish error:", err)
        return
    }

    fmt.Println(result)
}