package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			DisableSSL:       aws.Bool(true),
			S3ForcePathStyle: aws.Bool(true),
			Credentials:      credentials.NewStaticCredentials("user", "password", ""),
			Region:           aws.String("us-east-1"),
			Endpoint:         aws.String("nginx"),
		},
	}))

	svc := s3.New(sess)
	objs, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String("test"),
	})

	if err != nil {
		log.Fatal(err)
	}
	for _, obj := range objs.Contents {
		log.Println(*obj.Key)
	}
}
