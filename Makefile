user:
	protoc -I/usr/local/include  -I. \
		-I=$(GOPATH)/src -I=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=micro:. \
		users/server/proto/users.proto

user-gw:
	# grpc-gateway have to use grpc plugins so here we have to build users.proto again with grpc, too.
	protoc -I/usr/local/include  -I. \
		-I=$(GOPATH)/src -I=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		users/gateway/proto/users.proto

	protoc -I/usr/local/include -I. \
		-I=$(GOPATH)/src \
		-I=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		users/gateway/proto/users.proto

user-swagger:
	protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:. \
		users/gateway/proto/users.proto