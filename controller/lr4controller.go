package controller

import (
	"main/dto"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Lab4Controller struct {
	DB *sqlx.DB
}

func CreateLab4Controller(db *sqlx.DB) *Lab4Controller {
	return &Lab4Controller{DB: db}
}

// GetProducts godoc
// @Summary      Get all products
// @Description  Returns a list of all products
// @Tags         products
// @Produce      json
// @Success      200  {array}   dto.ProductDTO
// @Failure      404  {object}  dto.ServerErrorResponse
// @Failure      500  {object}  dto.ServerErrorResponse
// @Router       /products [get]
func (lc *Lab4Controller) GetProducts(context *gin.Context) {
	rows, err := lc.DB.Query("SELECT * FROM product")
	if err != nil {

		context.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	var products []dto.ProductDTO
	for rows.Next() {
		var product dto.ProductDTO
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt)
		if err != nil {
			context.JSON(500, dto.ServerErrorResponse{
				Error: err.Error(),
			})
			return
		}

		products = append(products, product)
	}

	if len(products) == 0 {
		context.JSON(404, dto.ServerErrorResponse{
			Error: "Not Found",
		})
		return
	}

	context.JSON(200, products)
}

// GetProduct godoc
// @Summary      Get product by ID
// @Description  Returns a single product by its ID
// @Tags         products
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  dto.ProductDTO
// @Failure      404  {object}  dto.ServerErrorResponse
// @Router       /products/{id} [get]
func (lc *Lab4Controller) GetProduct(context *gin.Context) {
	id := context.Param("id")

	var product dto.ProductDTO

	err := lc.DB.QueryRow(
		"SELECT * FROM product where id = $1",
		id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt)
	if err != nil {
		context.JSON(404, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	context.JSON(200, product)
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Creates a new product with the provided data
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      dto.CreateProductDTO  true  "Product data"
// @Success      201      {object}  dto.CreateProductDTO
// @Failure      400      {object}  dto.ServerErrorResponse
// @Failure      500      {object}  dto.ServerErrorResponse
// @Router       /products [post]
func (lc *Lab4Controller) CreateProduct(context *gin.Context) {
	var product dto.CreateProductDTO

	if err := context.ShouldBind(&product); err != nil {
		context.JSON(400, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	_, err := lc.DB.Exec(
		"INSERT INTO product(name, description , price, createdat) VALUES($1, $2, $3, $4)",
		product.Name, product.Description, product.Price, product.CreatedAt)
	if err != nil {
		context.JSON(500, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	context.JSON(201, product)
}

// UpdateProduct godoc
// @Summary      Update a product
// @Description  Updates an existing product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      dto.ProductDTO  true  "Updated product data"
// @Success      200      {object}  dto.ProductDTO
// @Failure      400      {object}  dto.ServerErrorResponse
// @Failure      404      {object}  dto.ServerErrorResponse
// @Failure      500      {object}  dto.ServerErrorResponse
// @Router       /products [put]
func (lc *Lab4Controller) UpdateProduct(context *gin.Context) {
	var product dto.ProductDTO

	if err := context.ShouldBind(&product); err != nil {
		context.JSON(400, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	res, err := lc.DB.Exec(
		"UPDATE product set name = $1, description = $2, price = $3, createdat = $4 where id = $5",
		product.Name, product.Description, product.Price, product.CreatedAt, product.ID)
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
			Error: "Id not found",
		})
		return
	}

	context.JSON(200, product)
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Deletes a product by ID
// @Tags         products
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      204  {object}  map[string]interface{}
// @Failure      500  {object}  dto.ServerErrorResponse
// @Failure      404  {object}  dto.ServerErrorResponse
// @Router       /products/{id} [delete]
func (lc *Lab4Controller) DeleteProduct(context *gin.Context) {
	id := context.Param("id")

	res, err := lc.DB.Exec(
		"delete from product where id = $1", id)
	if err != nil {
		context.JSON(500, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	deletedRowsNum, err := res.RowsAffected()
	if err != nil {
		context.JSON(500, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if deletedRowsNum == 0 {
		context.JSON(404, dto.ServerErrorResponse{
			Error: "Not found",
		})
		return
	}
	context.JSON(204, gin.H{})
}
