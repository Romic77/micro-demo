module micro-demo

go 1.15

require (
	github.com/asim/go-micro/cmd/protoc-gen-micro/v3 v3.7.0 // indirect
	github.com/asim/go-micro/plugins/registry/consul/v3 v3.7.0 // indirect
	github.com/asim/go-micro/v3 v3.7.1
	github.com/gin-gonic/gin v1.7.7 // indirect
	github.com/go-git/go-git/v5 v5.4.2 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/hashicorp/consul/api v1.12.0 // indirect
	github.com/micro/micro/v3 v3.10.1
	google.golang.org/protobuf v1.28.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.27.1
