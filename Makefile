.PHONY: proto
proto:
	@echo Make sure you have locally installed:
	@echo   protoc          v3.11.4
	@echo   protoc-gen-go    v1.4.2
	@echo   grpc-gateway    v1.14.6
	@sleep 2
	mkdir -p gen/go
	protoc --go_out=plugins=grpc:./gen/go --grpc-gateway_out=allow_patch_feature=false,allow_delete_body=true,logtostderr=true:./gen/go -Iproto svctwo/svctwo.proto
	protoc --go_out=plugins=grpc:./gen/go --grpc-gateway_out=allow_patch_feature=false,allow_delete_body=true,logtostderr=true:./gen/go -Iproto svcone/svcone.proto
	rm -rf ./svc
	mv gen/go/github.com/sgtsquiggs/grpc_gateway_bug/* ./

build: proto
	go build .
