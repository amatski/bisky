GOARCH=amd64 GOOS=linux go build compiler_lambda.go 
GOARCH=amd64 GOOS=linux go build judge_lambda.go && docker build -f Dockerfile .