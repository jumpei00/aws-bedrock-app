package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jumpei00/aws-bedrock-app/lambda/handler"
)

func main() {
	handler, err := handler.NewPostConversationHandler("conversations")
	if err != nil {
		log.Fatalf("failed to create handler: %+v", err)
	}
	lambda.Start(handler.Handle)
}
