package main

import (
	"context"
	"encoding/json"
	"os"
	"time"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Body struct {
	message string
	region  string
	time    int64
}

// time format "01/02/2006 03:04:05 PM (MST)"
func handleRequest(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	body := map[string]interface{}{
		"message": "Learn AWS the way you want",
		"region":  os.Getenv("AWS_REGION"),
		"time":    time.Now().Format(time.Kitchen),
	}

	jsonBody, jsonErr := json.Marshal(body)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	return events.APIGatewayV2HTTPResponse{
		Body:       string(jsonBody),
		StatusCode: 200,
		Headers:    headers,
	}, jsonErr
}

func main() {
	lambda.Start(handleRequest)
}
