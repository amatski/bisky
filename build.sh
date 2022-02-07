GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build compiler_lambda.go 
GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build judge_lambda.go && docker build -f Dockerfile .
