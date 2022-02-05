package main

import (
	"github.com/amatski/bisky/compiler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	r := compiler.NewRequestHandler()
	lambda.Start(r.HandleCompileRequest)
}
