test:
	go test -v ./judge/... ./compiler/...

lint: 
	golangci-lint run compiler/ judge/
   
proto:
	protoc --go_out=. --twirp_out=. rpc/bisky.proto

bisky:
	go run twirp/server.go
