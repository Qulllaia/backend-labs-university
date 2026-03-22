package lr6

import (
	"main/controller/lr6"
	"main/model"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterRouterForLR6(router *gin.RouterGroup, db *sqlx.DB) {
	orderModel := model.CreateOrderModel(db);
	orderService := lr6.OrderServiceCreate(orderModel);
	orderController := lr6.CreateOrderControllerLR6(orderService)

	router.GET("/order/:id", orderController.GetOrder)

	router.POST("/orders", orderController.CreateOrder)

}
