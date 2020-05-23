package repository

import (
	"github.com/jinzhu/gorm"
	"shopping/order/model"
)

type Repository interface {
	Create(*model.Order) error
	Find(string) (*model.Order, error)
	Update(*model.Order) (model.Order, error)
}

type Order struct {
	Db *gorm.DB
}

func (this *Order) Create(order *model.Order) error {
	if err := this.Db.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func (this *Order) Find(orderId string) (*model.Order, error) {
	order := &model.Order{}
	order.OrderId = orderId
	if err := this.Db.First(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (this *Order) Update(order *model.Order) (*model.Order, error) {
	if err := this.Db.Model(&order).Updates(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
