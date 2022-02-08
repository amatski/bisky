package main

import (
	"context"
	"net/http"

	"github.com/amatski/bisky/algotrainer/rpc/bisky"

	"github.com/rs/cors"
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
	handler := cors.AllowAll().Handler(twirpHandler)
	mux := http.NewServeMux()
	mux.Handle(twirpHandler.PathPrefix(), handler)
	http.ListenAndServe(":8080", mux)
}
