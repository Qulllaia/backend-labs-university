package dto

type ProductDTO struct {
	ID          int64   `db:"id" form:"id"`
	Name        string  `db:"name" form:"name" binding:"required,min=3"`
	Description *string `db:"description" form:"description" binding:"max=200"`
	Price       int64   `db:"price" form:"price" binding:"required,min=100,max=1000"`
	CreatedAt   *string `db:"createdat" form:"createdat"`
}

type CreateProductDTO struct {
	Name        string  `db:"name" form:"name" binding:"required,min=3"`
	Description *string `db:"description" form:"description" binding:"max=200"`
	Price       int64   `db:"price" form:"price" binding:"required,min=100,max=1000"`
	CreatedAt   *string `db:"createdat" form:"createdat"`
}

type ServerErrorResponse struct {
	Error string `json:"error"`
}
