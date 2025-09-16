package database

import (
	"context"
	"database/sql"
	"fmt"
	orderRoDataStore "github.com/devlibx/go-template-project/pkg/infra/database/mysql/user/ro"
	ordersDataStore "github.com/devlibx/go-template-project/pkg/infra/database/mysql/user/rw"
	"github.com/devlibx/gox-base/v2"
	"github.com/devlibx/gox-base/v2/errors"
	"go.uber.org/fx"
	"time"
)

// ConfigProvider interface defines methods to get database configuration
type ConfigProvider interface {
	SetupDefault()
	GetDatabase() string
	GetHost() string
	GetPort() int
	GetUser() string
	GetPassword() string
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
	GetConnMaxLifetimeInSec() int
	GetConnMaxIdleTimeInSec() int
}

type DbConnections struct {
	OrdersSqlDbConnection  *sql.DB
	OrderRoSqlDbConnection *sql.DB
}

var Provider = fx.Options(

	// Build all SQL connections here
	fx.Provide(func(ordersDataStoreCfg *ordersDataStore.MySqlConfig, orderRoDataStoreCfg *orderRoDataStore.MySqlConfig) (*DbConnections, error) {
		ordersSqlDbConnection, err := buildDatabaseConnection(ordersDataStoreCfg)
		if err != nil {
			return nil, err
		}
		orderRoSqlDbConnection, err := buildDatabaseConnection(orderRoDataStoreCfg)
		if err != nil {
			return nil, err
		}
		return &DbConnections{
			OrdersSqlDbConnection:  ordersSqlDbConnection,
			OrderRoSqlDbConnection: orderRoSqlDbConnection,
		}, nil
	}),

	// Build specific querier, and queries (e.g. RW connections)
	fx.Provide(func(cf gox.CrossFunction, dbConnections *DbConnections) (ordersDataStore.Querier, *ordersDataStore.Queries, error) {
		q, err := ordersDataStore.Prepare(context.Background(), dbConnections.OrdersSqlDbConnection)
		return q, q, err
	}),

	// Build specific querier, and queries (e.g. RO connections)
	fx.Provide(func(cf gox.CrossFunction, dbConnections *DbConnections) (orderRoDataStore.Querier, *orderRoDataStore.Queries, error) {
		q, err := orderRoDataStore.Prepare(context.Background(), dbConnections.OrdersSqlDbConnection)
		return q, q, err
	}),
)

func buildDatabaseConnection(configProvider ConfigProvider) (*sql.DB, error) {
	// Setup default values if missing
	configProvider.SetupDefault()

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		configProvider.GetUser(), configProvider.GetPassword(), configProvider.GetHost(), configProvider.GetPort(), configProvider.GetDatabase())
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, errors.Wrap(err, "error in connecting to database - failed to call sql.Open: database=[%s]", configProvider.GetDatabase())
	}

	// Connection configurations
	db.SetMaxOpenConns(configProvider.GetMaxOpenConnection())
	db.SetMaxIdleConns(configProvider.GetMaxIdleConnection())
	db.SetConnMaxLifetime(time.Duration(configProvider.GetConnMaxLifetimeInSec()) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(configProvider.GetConnMaxIdleTimeInSec()) * time.Second)

	return db, err
}
