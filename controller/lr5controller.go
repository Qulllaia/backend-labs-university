package controller

import (
	"fmt"
	"strconv"

	"main/dto"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var validStatuses = map[string]bool{
	"CREATED":    true,
	"DELIVERING": true,
	"DONE":       true,
}

var validSortFields = map[string]bool{
	"id": true, "productid": true, "status": true,
}

type OrderController struct {
	DB *sqlx.DB
}

func CreateOrderController(db *sqlx.DB) *OrderController {
	return &OrderController{DB: db}
}

// GetOrders godoc
// @Summary      Get all orders
// @Description  Returns a list of all orders
// @Tags         orders
// @Produce      json
// @Success      200  {array}   dto.OrderDTO
// @Failure      404  {object}  dto.ServerErrorResponse
// @Failure      500  {object}  dto.ServerErrorResponse
// @Router       /orders [get]
func (oc *OrderController) GetOrders(context *gin.Context) {
	productId := context.Param("productid")
	page := context.Query("page")
	pageSize := context.Query("pageSize")
	sortField := context.Query("sortField")

	query := "SELECT id, productid, status FROM \"order\" WHERE 1=1"
	params := map[string]interface{}{}

	if productId != "" {
		query += " AND productid = :productid"
		params["productid"] = productId
	}

	if sortField != "" {
		if !validSortFields[sortField] {
			context.JSON(400, dto.ServerErrorResponse{Error: "invalid sort field"})
			return
		}
		query += fmt.Sprintf("ORDER BY %s ASC ", sortField)
	}

	if page != "" && pageSize != "" {
		intPage, _ := strconv.Atoi(page)
		intPageSize, _ := strconv.Atoi(pageSize)

		if intPageSize > 0 && intPageSize <= 100 {
			query += " LIMIT :limit OFFSET :offset"
			params["limit"] = intPageSize
			params["offset"] = (intPage - 1) * intPageSize
		}
	}

	rows, err := oc.DB.NamedQuery(query, params)
	if err != nil {
		context.JSON(500, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}
	defer rows.Close()

	var orders []dto.OrderDTO
	for rows.Next() {
		var order dto.OrderDTO
		err := rows.Scan(&order.ID, &order.ProductID, &order.Status)
		if err != nil {
			context.JSON(500, dto.ServerErrorResponse{
				Error: err.Error(),
			})
			return
		}
		orders = append(orders, order)
	}

	if len(orders) == 0 {
		context.JSON(404, dto.ServerErrorResponse{
			Error: "Not Found",
		})
		return
	}

	context.JSON(200, orders)
}

// GetOrdersWithStatus godoc
// @Summary      Get orders by product ID and status
// @Description  Returns orders filtered by product ID and status
// @Tags         orders
// @Produce      json
// @Param        productId path int true "Product ID"
// @Param        status path string true "Order status" Enums(CREATED, DELIVERING, DONE)
// @Success      200  {array}   dto.OrderDTO
// @Failure      404  {object}  dto.ServerErrorResponse
// @Failure      500  {object}  dto.ServerErrorResponse
// @Router       /orders/product/{productId}/status/{status} [get]
func (oc *OrderController) GetOrdersWithStatus(context *gin.Context) {
	var OrderAndStatus dto.URIOrderWithProductIDAndStatus

	if err := context.ShouldBindUri(&OrderAndStatus); err != nil {
		context.JSON(500, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	rows, err := oc.DB.Query("SELECT id, productid, status FROM \"order\" WHERE productid = $1 AND status = $2", OrderAndStatus.ProductID, OrderAndStatus.Status)
	if err != nil {
		context.JSON(500, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}
	defer rows.Close()

	var orders []dto.OrderDTO
	for rows.Next() {
		var order dto.OrderDTO
		err := rows.Scan(&order.ID, &order.ProductID, &order.Status)
		if err != nil {
			context.JSON(500, dto.ServerErrorResponse{
				Error: err.Error(),
			})
			return
		}
		orders = append(orders, order)
	}

	if len(orders) == 0 {
		context.JSON(404, dto.ServerErrorResponse{
			Error: "Not Found",
		})
		return
	}

	context.JSON(200, orders)
}

// GetOrder godoc
// @Summary      Get order by ID
// @Description  Returns a single order by its ID
// @Tags         orders
// @Produce      json
// @Param        id   path      int  true  "Order ID"
// @Success      200  {object}  dto.OrderDTO
// @Failure      404  {object}  dto.ServerErrorResponse
// @Router       /orders/{id} [get]
func (oc *OrderController) GetOrder(context *gin.Context) {
	id := context.Param("id")

	var order dto.OrderDTO
	err := oc.DB.QueryRow(
		"SELECT id, productid, status FROM \"order\" WHERE id = $1",
		id).Scan(&order.ID, &order.ProductID, &order.Status)
	if err != nil {
		context.JSON(404, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	context.JSON(200, order)
}

// CreateOrder godoc
// @Summary      Create a new order
// @Description  Creates a new order with the provided data
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        order  body      dto.CreateOrderDTO  true  "Order data"
// @Success      201    {object}  dto.OrderDTO
// @Failure      400    {object}  dto.ServerErrorResponse
// @Failure      500    {object}  dto.ServerErrorResponse
// @Router       /orders [post]
func (oc *OrderController) CreateOrder(context *gin.Context) {
	var order dto.CreateOrderDTO

	if err := context.ShouldBindJSON(&order); err != nil {
		context.JSON(400, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if !validStatuses[order.Status] {
		context.JSON(400, dto.ServerErrorResponse{
			Error: "invalid status",
		})
		return
	}

	var exists bool
	err := oc.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM product WHERE id = $1)", order.ProductID).Scan(&exists)
	if err != nil {
		context.JSON(500, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if !exists {
		context.JSON(400, dto.ServerErrorResponse{
			Error: "product not found",
		})
		return
	}

	var id int64
	err = oc.DB.QueryRow(
		"INSERT INTO \"order\" (productid, status) VALUES($1, $2) RETURNING id",
		order.ProductID, order.Status).Scan(&id)
	if err != nil {
		context.JSON(500, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	createdOrder := dto.OrderDTO{
		ID:        id,
		ProductID: order.ProductID,
		Status:    order.Status,
	}

	context.JSON(201, createdOrder)
}

// pdateOrder godoc
// @Summary      Update an order
// @Description  Updates an existing order by ID
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id     path      int               true  "Order ID"
// @Param        order  body      dto.UpdateOrderDTO  true  "Updated order data"
// @Success      200    {object}  dto.OrderDTO
// @Failure      400    {object}  dto.ServerErrorResponse
// @Failure      404    {object}  dto.ServerErrorResponse
// @Failure      500    {object}  dto.ServerErrorResponse
// @Router       /orders/{id} [put]
func (oc *OrderController) UpdateOrder(context *gin.Context) {
	id := context.Param("id")

	var updateData dto.UpdateOrderDTO
	if err := context.ShouldBindJSON(&updateData); err != nil {
		context.JSON(400, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if !validStatuses[updateData.Status] {
		context.JSON(400, dto.ServerErrorResponse{
			Error: "invalid status",
		})
		return
	}

	result, err := oc.DB.Exec(
		"UPDATE \"order\" SET status = $1 WHERE id = $2",
		updateData.Status, id)
	if err != nil {
		context.JSON(500, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		context.JSON(500, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		context.JSON(404, dto.ServerErrorResponse{
			Error: "sql: no rows in result set",
		})
		return
	}

	var updatedOrder dto.OrderDTO
	err = oc.DB.QueryRow(
		"SELECT id, productid, status FROM \"order\" WHERE id = $1",
		id).Scan(&updatedOrder.ID, &updatedOrder.ProductID, &updatedOrder.Status)
	if err != nil {
		context.JSON(500, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	context.JSON(200, updatedOrder)
}

// DeleteOrder godoc
// @Summary      Delete an order
// @Description  Deletes an order by ID
// @Tags         orders
// @Produce      json
// @Param        id   path      int  true  "Order ID"
// @Success      204  {object}  map[string]interface{}
// @Failure      404  {object}  dto.ServerErrorResponse
// @Failure      500  {object}  dto.ServerErrorResponse
// @Router       /orders/{id} [delete]
func (oc *OrderController) DeleteOrder(context *gin.Context) {
	id := context.Param("id")

	res, err := oc.DB.Exec("DELETE FROM \"order\" WHERE id = $1", id)
	if err != nil {
		context.JSON(500, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		context.JSON(500, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		context.JSON(404, dto.ServerErrorResponse{
			Error: "sql: no rows in result set",
		})
		return
	}

	context.JSON(204, gin.H{})
}
