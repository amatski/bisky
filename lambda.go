package main

import (
	"flag"

	"github.com/amatski/bisky/compiler"
	"github.com/amatski/bisky/judge"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	target string
)

func main() {
	flag.StringVar(&target, "target", "judge", "specify server target (compiler or judge)")
	if target == "compiler" {
		r := compiler.NewRequestHandler()
		lambda.Start(r.HandleCompileRequest)
	} else if target == "judge" {
		r := judge.NewRequestHandler()
		lambda.Start(r.JudgeSolution)
	}
}
