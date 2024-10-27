package grpc

import "githib.com/isurucuma/go-grpc/microservices/order/internal/ports"

type Adapter struct {
	api  ports.APIPort
	port int
	//order.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}
