package repository

import (
	"github.com/jinzhu/gorm"
	"shopping/product/model"

	product1 "shopping/product/proto/product"
)

type Repository interface {
	Find(id int32) (*model.Product, error)
	Create(*model.Product) error
	Update(*model.Product) error
	FindByField(string, string, string) (*model.Product, error)
	FindProducts(string) ([]*model.Product, error)
	FindDetail(id int32) (*model.Product, error)
}

type Product struct {
	Db *gorm.DB
}

func (this *Product) Find(id uint32) (*model.Product, error) {
	product := &model.Product{}
	product.ID = uint(id)
	if err := this.Db.First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (this *Product) Create(product *model.Product) error {
	if err := this.Db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (this *Product) Update(product *model.Product) error {
	if err := this.Db.Model(product).Where("id=?", &product.ID).Update("number", &product.Number).Error; err != nil {
		return err
	}
	return nil
}

func (this *Product) FindByField(key string, value string, fields string) (*model.Product, error) {
	if len(fields) == 0 {
		fields = "*"
	}
	product := &model.Product{}
	if err := this.Db.Select(fields).Where(key+"=?", value).First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (this *Product) FindProducts(name string) ([]*product1.Product, error) {
	var products []*product1.Product
	if err := this.Db.Where("name like ?", "%"+name+"%").Limit(10).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (this *Product) FindDetail(id uint32) (*product1.Product, error) {
	var product = &product1.Product{}
	if err := this.Db.Where("id=?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
