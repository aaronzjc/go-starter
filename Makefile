GO111MODULE=on
GOPROXY=https://goproxy.cn,direct
protodir="./internal/pb"

setup:
	@GOPROXY=$(GOPROXY) CGO_ENABLED=0 GOOS=linux GOARCH=amd64
build_api:
	 go build -o ./bin/api ./cmd/api/main.go
build_grpc:
	@GOPROXY=$(GOPROXY) CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/grpc ./cmd/grpc/main.go
gen_proto:
	protoc --go_out=. --go_opt=paths=source_relative $(protodir)/*.proto
gen_grpc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(protodir)/*.proto
clean:
	-rm ./bin/*

.PHONY: build_api build_grpc gen_proto gen_grpc clean