package lr4

import (
	"main/controller"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterRouterForLR4(router *gin.RouterGroup, db *sqlx.DB) {
	lab4controller := controller.CreateLab4Controller(db)
	router.GET("/products", lab4controller.GetProducts)

	router.GET("/products/:id", lab4controller.GetProduct)

	router.POST("/products", lab4controller.CreateProduct)

	router.PUT("/products", lab4controller.UpdateProduct)

	router.DELETE("/products/:id", lab4controller.DeleteProduct)
}
