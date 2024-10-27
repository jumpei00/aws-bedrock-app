package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jumpei00/aws-bedrock-app/lambda/handler"
)

func main() {
	handler, err := handler.NewGetConversationHandler("conversations")
	if err != nil {
		log.Fatalf("failed to create handler: %+v", err)
	}
	lambda.Start(handler.Handle)
}
