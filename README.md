# readr-micro
readr-micro is an experimental gRPC revision of [readr-restful](https://github.com/readr-media/readr-restful).
Adopting go-grpc, grpc-gateway and postgresql

It only revised part of members now.

## gRPC service

### Build

```bash
make build-pb
```

### Run
Currently run under **mdns** mode
```bash
go run server/*.go --server_address=localhost:50051 --registry=mdns
```

## RESTful reverse proxy

### build
```bash
make build-gw
```

### Run
```bash
go run gateway/main.go
```
This will run a RESTful server at port `8080`. It could be requested with `http://localhost:8080/members/{id}`