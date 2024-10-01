package main

import (
	"challenge/internal/handlers"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handlers.HandleS3Event)
}
