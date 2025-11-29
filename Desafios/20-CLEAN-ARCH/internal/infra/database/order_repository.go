package database

import (
	"context"
	"database/sql"

	"github.com/Frank-Macedo/20-cleanArch/internal/db"
	"github.com/Frank-Macedo/20-cleanArch/internal/entity"
)

type OrderRepository struct {
	Db      *sql.DB
	Queries *db.Queries
}

func NewOrderRepository(dbConn *sql.DB) *OrderRepository {
	return &OrderRepository{Db: dbConn, Queries: db.New(dbConn)}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (c *OrderRepository) FindAll() ([]entity.Order, error) {
	dbOrders, err := c.Queries.Getallorders(context.Background())
	if err != nil {
		return nil, err
	}
	orders := []entity.Order{}

	for _, dbOrder := range dbOrders {
		order := entity.Order{
			ID:         dbOrder.ID,
			Price:      dbOrder.Price,
			Tax:        dbOrder.Tax,
			FinalPrice: dbOrder.FinalPrice,
		}
		orders = append(orders, order)
	}
	return orders, nil
}
