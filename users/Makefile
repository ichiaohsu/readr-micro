build-pb:
	protoc -I/usr/local/include  -I. \
		-I=$(GOPATH)/src -I=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=micro:. \
		server/proto/users.proto

build-gw:

	protoc -I/usr/local/include  -I. \
		-I=$(GOPATH)/src -I=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		gateway/proto/users.proto

	protoc -I/usr/local/include -I. \
		-I=$(GOPATH)/src \
		-I=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		gateway/proto/users.proto