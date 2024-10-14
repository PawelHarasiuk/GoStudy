package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	mess string
}

func handleRequest() (string, error) {
	fmt.Println("hi")
	return "Hello from Lambda", nil
}

func main() {
	lambda.Start(handleRequest)
}
