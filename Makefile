test:
	go test -v ./judge/... ./compiler/...

lint: 
	golangci-lint run compiler/ judge/
