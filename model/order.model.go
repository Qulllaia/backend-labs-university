package model

import (
	"main/dto"

	"github.com/jmoiron/sqlx"
)

type IOrderModel interface {
	GetOrder(id string, order *dto.OrderDTO) error
	CheckOrderExistence(id int64, exists *bool) error
	CreateOrder(order dto.CreateOrderDTO) (int64, error)
}

type OrderModel struct {
	DB *sqlx.DB
}

func CreateOrderModel(db *sqlx.DB) *OrderModel {
	return &OrderModel{DB: db}
}


func (om *OrderModel) GetOrder(id string, order *dto.OrderDTO) error {
	return om.DB.QueryRow(
		"SELECT id, productid, status FROM \"order\" WHERE id = $1",
		id).Scan(&order.ID, &order.ProductID, &order.Status)
	
}

func (om *OrderModel) CheckOrderExistence(id int64, exists *bool) error {
	return om.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM product WHERE id = $1)", id).Scan(exists)
}

func (om *OrderModel) CreateOrder(order dto.CreateOrderDTO) (int64, error) {

	var id int64
	err := om.DB.QueryRow(
		"INSERT INTO \"order\" (productid, status) VALUES($1, $2) RETURNING id",
		order.ProductID, order.Status).Scan(&id)
	return id, err
}

