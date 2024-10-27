package domain

import "time"

type OrderItem struct {
	ProductCode string  `json:"product_code"`
	Quantity    int32   `json:"quantity"`
	UnitPrice   float32 `json:"unit_price"`
}

type Order struct {
	ID         int64       `json:"id"`
	CustomerId int64       `json:"customer_id"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  int64       `json:"created_at"`
}

func NewOrder(customerId int64, orderItems []OrderItem) Order {
	return Order{
		CustomerId: customerId,
		OrderItems: orderItems,
		Status:     "pending",
		CreatedAt:  time.Now().Unix(),
	}
}
