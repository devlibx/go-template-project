package base

import (
	ordersDataStore "github.com/devlibx/go-template-project/pkg/infra/database/mysql/user/rw"
	goxHttpApi "github.com/devlibx/gox-http/v4/api"
)

type ApplicationContext struct {
	GoxHttpContext  goxHttpApi.GoxHttpContext
	OrdersDataStore ordersDataStore.Querier
}
