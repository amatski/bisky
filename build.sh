GOARCH=amd64 GOOS=linux go build ./... && docker build -f Dockerfile .
