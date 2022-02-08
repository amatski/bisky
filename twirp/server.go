package main

import (
	"context"
	"log"
	"net/http"

	"github.com/amatski/bisky/algotrainer/rpc/bisky"
	"github.com/amatski/bisky/judge"
	"github.com/amatski/bisky/judge/codegen"
	"github.com/amatski/bisky/judge/problem"

	"github.com/rs/cors"
)

type BiskyServer struct {
}

func (s *BiskyServer) Judge(ctx context.Context, req *bisky.JudgeRequest) (*bisky.JudgeResponse, error) {

	handler := judge.RequestHandler{}

	tests := []*problem.TestCase{}

	for _, test := range req.TestCases {
		tests = append(tests, &problem.TestCase{
			Input:          test.Input,
			ExpectedOutput: test.ExpectedOutput,
		})
	}

	log.Println("calling res from judge", req, req.TestCases)
	res, err := handler.JudgeSolution(judge.JudgeRequest{
		Language:    req.Language,
		EncodedCode: req.EncodedCode,
		Problem:     "two_sum",
		OutputType:  codegen.Integers,
		TestCases:   tests,
	})

	log.Println(err, res, "res from judge")

	if err != nil {
		return nil, err
	}

	ress := []*bisky.TestCaseResult{}
	for _, v := range res.Results {
		ress = append(ress, &bisky.TestCaseResult{
			Stdout:  v.Stdout,
			Passed:  v.Passed,
			Elapsed: "",
		})

	}

	log.Println(ress)

	return &bisky.JudgeResponse{
		Results: ress,
	}, nil
}

func main() {
	twirpHandler := bisky.NewBiskyServer(&BiskyServer{})
	handler := cors.AllowAll().Handler(twirpHandler)
	mux := http.NewServeMux()
	mux.Handle(twirpHandler.PathPrefix(), handler)
	http.ListenAndServe(":8080", mux)
}
