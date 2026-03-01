package dto

type OrderDTO struct {
	ID        int64  `json:"id" db:"id"`
	ProductID int64  `json:"productId" db:"productid"`
	Status    string `json:"status" db:"status" enums:"CREATED,DELIVERING,DONE"`
}

type CreateOrderDTO struct {
	ProductID int64  `json:"productId" binding:"required"`
	Status    string `json:"status" binding:"required" enums:"CREATED,DELIVERING,DONE"`
}

type UpdateOrderDTO struct {
	Status string `json:"status" binding:"required" enums:"CREATED,DELIVERING,DONE"`
}

type URIOrderWithProductIDAndStatus struct {
	ProductID int64  `uri:"productid" binding:"required,min=0"`
	Status    string `uri:"status" binding:"required" enums:"CREATED,DELIVERING,DONE"`
}
