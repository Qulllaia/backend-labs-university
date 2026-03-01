package lr5

import (
	"main/controller"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterRouterForLR5(router *gin.RouterGroup, db *sqlx.DB) {
	orderController := controller.CreateOrderController(db)
	router.GET("/orders", orderController.GetOrders)

	router.GET("/orders/:productid", orderController.GetOrders)

	router.GET("/orders/product/:productid/status/:status", orderController.GetOrdersWithStatus)

	router.GET("/order/:id", orderController.GetOrder)

	router.POST("/orders", orderController.CreateOrder)

	router.PUT("/orders", orderController.UpdateOrder)

	router.DELETE("/orders/:id", orderController.DeleteOrder)
}
