package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := "World"

	if queryName, ok := request.QueryStringParameters["name"]; ok {
		name = queryName
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello %s", name),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
