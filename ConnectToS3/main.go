package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"log"
	"os"
)

func main() {
	bucket := "newsletter-bucket-go"
	key := "emails.txt"

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	svc := s3.NewFromConfig(cfg)

	downloader := manager.NewDownloader(svc)

	file, err := os.Create("/tmp/" + key)
	if err != nil {
		log.Fatalf("Error creating file", err)
	}

	_, err = downloader.Download(context.TODO(), file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	cont, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(cont))
	_, err = file.WriteString("hi\n")
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err.Error())
	}
	uploader := manager.NewUploader(svc)

	f, err := os.Open("/tmp/" + key)
	if err != nil {
		log.Fatal(err.Error())
	}
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        f,
		ContentType: aws.String("text/plain"),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("success: ", result)
}
