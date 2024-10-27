package grpc

import (
	"github.com/isurucuma/go-grpc/microservices/order/internal/ports"
	_go "github.com/isurucuma/go-grpc/microservices/protos/order/golang"
)

type Adapter struct {
	api  ports.APIPort
	port int
	_go.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}
