package models

import "time"

type Order struct {
	ID              int        `json:"id"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	Number          string     `json:"number"`
	CustomerName    string     `json:"customer_name"`
	Type            string     `json:"type"`
	TableNumber     *int       `json:"table_number,omitempty"`
	DeliveryAddress *string    `json:"delivery_address,omitempty"`
	TotalAmount     float64    `json:"total_amount"`
	Priority        int        `json:"priority"`
	Status          string     `json:"status"`
	ProcessedBy     *string    `json:"processed_by,omitempty"`
	CompletedAt     *time.Time `json:"completed_at,omitempty"`
}

type OrderItem struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	OrderID   int       `json:"order_id"`
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
}

type OrderStatusLog struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	OrderID   int       `json:"order_id"`
	Status    string    `json:"status"`
	ChangedBy time.Time `json:"changed_by"`
	ChangedAt time.Time `json:"changed_at"`
	Notes     *string   `json:"notes,omitempty"`
}

type CreateOrderRequest struct {
	CustomerName    string             `json:"customer_name"`
	OrderType       string             `json:"order_type"`
	TableNumber     *int               `json:"table_number,omitempty"`
	DeliveryAddress *string            `json:"delivery_address,omitempty"`
	Items           []OrderItemRequest `json:"items"`
}

type OrderItemRequest struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type CreateOrderResponse struct {
	OrderNumber string  `json:"order_number"`
	Status      string  `json:"status"`
	TotalAmount float64 `json:"total_amount"`
}

type OrderMessage struct {
	OrderNumber     string             `json:"order_number"`
	CustomerName    string             `json:"customer_name"`
	OrderType       string             `json:"order_type"`
	TableNumber     *int               `json:"table_number,omitempty"`
	DeliveryAddress *string            `json:"delivery_address,omitempty"`
	Items           []OrderItemRequest `json:"items"`
	TotalAmount     float64            `json:"total_amount"`
	Priority        int                `json:"priority"`
}
