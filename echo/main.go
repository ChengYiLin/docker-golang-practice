package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

var echoLambda *echoadapter.EchoLambda

// Hello
type Hello struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

func init() {
	log.Printf("Start Golang Echo")
	e := echo.New()

	hello := &Hello{
		Name:  "JohnCiner",
		Email: "johnciner@gmail.com",
	}
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, hello)
	})

	echoLambda = echoadapter.New(e)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, request)
}

func main() {
	lambda.Start(handler)
}
