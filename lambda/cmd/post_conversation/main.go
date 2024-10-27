package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler() {
	fmt.Println("Hello, World!")
}

func main() {
	lambda.Start(handler)
}
