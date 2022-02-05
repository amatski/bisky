test:
	go test ./judge/... ./compiler/...

lint: 
	golangci-lint run
