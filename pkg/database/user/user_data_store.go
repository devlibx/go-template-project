package user

import (
	"context"
	ordersDataStore "github.com/devlibx/go-template-project/pkg/infra/database/mysql/user/rw"
	orderRoDataStore "github.com/devlibx/go-template-project/pkg/infra/database/mysql/user/ro"
	"github.com/devlibx/gox-base/v2"
)

// UserDataStore interface defines all operations for users (both read and write)
type UserDataStore interface {
	// Write operations (use RW connection)
	CreateUser(ctx context.Context, arg CreateUserRequest) error
	UpdateUser(ctx context.Context, userID string, arg UpdateUserRequest) error
	DeleteUser(ctx context.Context, userID string) error
	
	// Read operations (use RO connection)
	GetUserByID(ctx context.Context, userID string) (*User, error)
	GetAllUsers(ctx context.Context) ([]*User, error)
}

// OrderDataStore interface defines operations for orders (keeping existing functionality)
type OrderDataStore interface {
	CreateOrder(ctx context.Context, arg CreateOrderRequest) error
	GetAllOrders(ctx context.Context) ([]*Order, error)
	GetOrderByID(ctx context.Context, orderID string) (*Order, error)
}

// userDataStoreImpl implements all user operations with both RO and RW connections
type userDataStoreImpl struct {
	gox.CrossFunction
	// RW connection for write operations
	rwQuerier ordersDataStore.Querier
	rwQueries *ordersDataStore.Queries
	// RO connection for read operations
	roQuerier orderRoDataStore.Querier
	roQueries *orderRoDataStore.Queries
}

// orderDataStoreImpl implements operations for orders (keeping existing functionality)
type orderDataStoreImpl struct {
	gox.CrossFunction
	querier ordersDataStore.Querier
	queries *ordersDataStore.Queries
}

// Write operations using RW connection
func (u *userDataStoreImpl) CreateUser(ctx context.Context, arg CreateUserRequest) error {
	// Using CreateOrder as underlying storage (demo purposes)
	return u.rwQuerier.CreateOrder(ctx, ordersDataStore.CreateOrderParams{
		OrderID:  arg.UserID,
		OrderQty: 1,        // Default value
		Amount:   arg.Email, // Using Amount field to store email
	})
}

func (u *userDataStoreImpl) UpdateUser(ctx context.Context, userID string, arg UpdateUserRequest) error {
	// Note: This would require an UPDATE query in the sqlc queries
	// For now, this is a placeholder implementation
	// You would need to add UpdateOrder or UpdateUser query to query.sql
	return nil // Placeholder
}

func (u *userDataStoreImpl) DeleteUser(ctx context.Context, userID string) error {
	// Note: This would require a DELETE query in the sqlc queries
	// For now, this is a placeholder implementation
	// You would need to add DeleteOrder or DeleteUser query to query.sql
	return nil // Placeholder
}

// Read operations using RO connection
func (u *userDataStoreImpl) GetUserByID(ctx context.Context, userID string) (*User, error) {
	if order, err := u.roQuerier.GetOrderByID(ctx, userID); err != nil {
		return nil, err
	} else {
		ret := &User{}
		ret.FromOrderRO(ctx, order)
		return ret, nil
	}
}

func (u *userDataStoreImpl) GetAllUsers(ctx context.Context) ([]*User, error) {
	if orders, err := u.roQuerier.GetAllOrders(ctx); err != nil {
		return nil, err
	} else {
		ret := make([]*User, len(orders))
		for i, order := range orders {
			ret[i] = &User{}
			ret[i].FromOrderRO(ctx, order)
		}
		return ret, nil
	}
}

// Order operations (existing functionality)
func (o *orderDataStoreImpl) CreateOrder(ctx context.Context, arg CreateOrderRequest) error {
	return o.querier.CreateOrder(ctx, ordersDataStore.CreateOrderParams{
		OrderID:  arg.OrderID,
		OrderQty: int32(arg.OrderQty),
		Amount:   arg.Amount,
	})
}

func (o *orderDataStoreImpl) GetAllOrders(ctx context.Context) ([]*Order, error) {
	if orders, err := o.querier.GetAllOrders(ctx); err != nil {
		return nil, err
	} else {
		ret := make([]*Order, len(orders))
		for i, order := range orders {
			ret[i] = &Order{}
			ret[i].FromOrder(ctx, order)
		}
		return ret, nil
	}
}

func (o *orderDataStoreImpl) GetOrderByID(ctx context.Context, orderID string) (*Order, error) {
	if order, err := o.querier.GetOrderByID(ctx, orderID); err != nil {
		return nil, err
	} else {
		ret := &Order{}
		ret.FromOrder(ctx, order)
		return ret, nil
	}
}

// Constructor functions
func NewUserDataStore(
	cf gox.CrossFunction, 
	rwQuerier ordersDataStore.Querier, 
	rwQueries *ordersDataStore.Queries,
	roQuerier orderRoDataStore.Querier, 
	roQueries *orderRoDataStore.Queries,
) UserDataStore {
	return &userDataStoreImpl{
		CrossFunction: cf,
		rwQuerier:     rwQuerier,
		rwQueries:     rwQueries,
		roQuerier:     roQuerier,
		roQueries:     roQueries,
	}
}

func NewOrderDataStore(cf gox.CrossFunction, querier ordersDataStore.Querier, queries *ordersDataStore.Queries) OrderDataStore {
	return &orderDataStoreImpl{
		CrossFunction: cf,
		querier:       querier,
		queries:       queries,
	}
}