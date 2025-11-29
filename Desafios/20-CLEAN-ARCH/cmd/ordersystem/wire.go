//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/Frank-Macedo/20-cleanArch/internal/entity"
	"github.com/Frank-Macedo/20-cleanArch/internal/event"
	"github.com/Frank-Macedo/20-cleanArch/internal/infra/database"
	"github.com/Frank-Macedo/20-cleanArch/internal/infra/web"
	"github.com/Frank-Macedo/20-cleanArch/internal/usecase"
	"github.com/Frank-Macedo/20-cleanArch/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewGetOrderUseCase(db *sql.DB) *usecase.GetOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewGetOrdersUseCase,
	)
	return &usecase.GetOrdersUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
