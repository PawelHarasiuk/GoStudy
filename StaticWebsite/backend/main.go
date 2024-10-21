package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Hello from lambda")
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		//Headers:           nil,
		Body: "Response body",
	}, nil
}

func main() {
	lambda.Start(handler)
}
