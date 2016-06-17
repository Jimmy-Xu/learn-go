package main

import (
	"fmt"
	"time"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func main() {
	accessKey := "AKIxxxxxxxxxxxxx"
	secretKey := "xxxxxxxxxxxxxxxx"
	region := "us-east-1"

	//new
	//svc := s3.New(session.New(&aws.Config{
	//	Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	//	Region: aws.String(region),
	//}))

	//old
	svc := s3.New(&aws.Config{
	       Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	       Region: &region,
	})

	r, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("image-tarball"),
		Key:    aws.String("test/private/nats.gz"),
	})
	url, err := r.Presign(5 * time.Minute)
	//url, err := r.Presign(7 * 24 * 60 * time.Minute)
	if err != nil {
		fmt.Println("error presigning request", err)
		return
	}
	fmt.Println("URL", url)
}
