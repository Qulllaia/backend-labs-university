package lr6

import (
	"main/dto"
	"main/model"
)

type IOrderService interface {
	GetOrder(id string) (*dto.OrderDTO, error)
	DoesOrderExists(id int64) bool 
	CreateOrder(dto.CreateOrderDTO) (*dto.OrderDTO, error)
}

type OrderService struct {
	OrderModel model.IOrderModel	
}

func OrderServiceCreate(orderModel model.IOrderModel) *OrderService {
	return &OrderService{OrderModel: orderModel}
}

func (oss OrderService) GetOrder(id string) (*dto.OrderDTO, error) {
	
	var order dto.OrderDTO
	if err := oss.OrderModel.GetOrder(id, &order); err != nil {
		return nil, err;
	}
	return &order, nil;
}


func (oss OrderService) DoesOrderExists(id int64) bool {
	var exists bool
	oss.OrderModel.CheckOrderExistence(id, &exists)
	return exists
} 
func (oss OrderService) CreateOrder(created dto.CreateOrderDTO) (*dto.OrderDTO, error)  {

	id, err := oss.OrderModel.CreateOrder(created)	
	if err != nil {
		return nil, err
	}
	
	createdOrder := dto.OrderDTO{
		ID:        id,
		ProductID: created.ProductID,
		Status:    created.Status,
	}

	return &createdOrder, nil
	
}
