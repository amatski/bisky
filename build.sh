GOARCH=amd64 GOOS=linux go build lambda.go && docker build -f Dockerfile .
