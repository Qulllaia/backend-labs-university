package lr6;

import (
	"main/dto"

	"github.com/gin-gonic/gin"
)


var validStatuses = map[string]bool{
	"CREATED":    true,
	"DELIVERING": true,
	"DONE":       true,
}

type OrderControllerLR6 struct {
	OrderService IOrderService
}

func CreateOrderControllerLR6(orderService IOrderService) *OrderControllerLR6 {
	return &OrderControllerLR6{OrderService: orderService}
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
func (oc *OrderControllerLR6) GetOrder(context *gin.Context) {
	id := context.Param("id")

	order, err := oc.OrderService.GetOrder(id);

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
func (oc *OrderControllerLR6) CreateOrder(context *gin.Context) {
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

	var exists bool = oc.OrderService.DoesOrderExists(order.ProductID) 

	if !exists {
		context.JSON(400, dto.ServerErrorResponse{
			Error: "product not found",
		})
		return
	}

	createdOrder, err := oc.OrderService.CreateOrder(order);

	if err != nil {
		context.JSON(500, dto.ServerErrorResponse{
			Error: err.Error(),
		})
		return
	}
	context.JSON(201, createdOrder)
}

