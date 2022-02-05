package main

import (
	"github.com/amatski/bisky/judge"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	r := judge.NewRequestHandler()
	lambda.Start(r.JudgeSolution)
}
