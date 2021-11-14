package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func helloWorld() {
	fmt.Printf("Hello World!")
}

func main() {
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return lambda.Start(helloWorld)
	// })
	lambda.Start(helloWorld)
}
