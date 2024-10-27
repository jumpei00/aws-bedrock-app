package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler() {
	fmt.Println("Hello, World!")
}

func main() {
	lambda.Start(Handler)
}
