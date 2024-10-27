package ports

import "githib.com/isurucuma/go-grpc/microservices/order/internal/application/core/domain"

type DBPort interface {
	Save(order *domain.Order) error
	Get(id string) (domain.Order, error)
}
