package database

import (
	"database/sql"

	"github.com/luizeduu/intensiveGo/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (repository *OrderRepository) Save(order *entity.Order) error {
	_, err := repository.Db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)", order.ID, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}

	return nil
}

func (repository *OrderRepository) GetTotalTransaction() (int, error) {

	var total int

	err := repository.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}
