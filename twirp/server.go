package main

import (
	"context"
	"net/http"

	"github.com/amatski/bisky/algotrainer/rpc/bisky"
)

type BiskyServer struct {
}

func (s *BiskyServer) Judge(ctx context.Context, req *bisky.JudgeRequest) (*bisky.JudgeResponse, error) {
	return &bisky.JudgeResponse{
		Results: []*bisky.TestCaseResult{
			{
				Stdout:  "no",
				Passed:  true,
				Elapsed: "0",
			},
		},
	}, nil
}

func main() {
	twirpHandler := bisky.NewBiskyServer(&BiskyServer{})
	mux := http.NewServeMux()
	// The generated code includes a method, PathPrefix(), which
	// can be used to mount your service on a mux.
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	http.ListenAndServe(":8080", mux)
}
