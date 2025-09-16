package user

import (
	"context"
	orderRoDataStore "github.com/devlibx/go-template-project/pkg/infra/database/mysql/user/ro"
	ordersDataStore "github.com/devlibx/go-template-project/pkg/infra/database/mysql/user/rw"
	"time"
)

// User represents a user in the system
type User struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserRequest represents a request to create a new user
type CreateUserRequest struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// UpdateUserRequest represents a request to update an existing user
type UpdateUserRequest struct {
	Email  string `json:"email,omitempty"`
	Name   string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
}

// FromOrder converts from sqlc generated Order type to domain User type
// Note: This assumes the orders table is being used for users (common in examples)
func (u *User) FromOrder(ctx context.Context, in *ordersDataStore.Order) *User {
	return &User{
		UserID:    in.OrderID,  // Using OrderID as UserID for demo
		Email:     in.Amount,   // Using Amount field as Email for demo
		Name:      "User Name", // Placeholder
		Status:    "active",    // Default status
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// FromOrderRO converts from sqlc generated RO Order type to domain User type
func (u *User) FromOrderRO(ctx context.Context, in *orderRoDataStore.Order) *User {
	return &User{
		UserID:    in.OrderID,  // Using OrderID as UserID for demo
		Email:     in.Amount,   // Using Amount field as Email for demo
		Name:      "User Name", // Placeholder
		Status:    "active",    // Default status
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// Order represents an order (keeping existing functionality)
type Order struct {
	OrderID   string    `json:"order_id"`
	OrderQty  int       `json:"order_qty"`
	Amount    string    `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateOrderRequest represents a request to create a new order
type CreateOrderRequest struct {
	OrderID  string `json:"order_id"`
	OrderQty int    `json:"order_qty"`
	Amount   string `json:"amount"`
}

// FromOrder converts from sqlc generated type to domain Order type
func (o *Order) FromOrder(ctx context.Context, in *ordersDataStore.Order) *Order {
	return &Order{
		OrderID:   in.OrderID,
		OrderQty:  int(in.OrderQty),
		Amount:    in.Amount,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
