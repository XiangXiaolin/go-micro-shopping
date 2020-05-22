package repository

import (
	"github.com/jinzhu/gorm"
	"shopping/user/model"
)

type Repository interface {
	Find(id int32) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) (*model.User, error)
	FindByField(string, string, string) (*model.User, error)
}

type User struct {
	Db *gorm.DB
}

func (this *User) Find(id uint32) (*model.User, error) {
	user := &model.User{}
	user.ID = uint(id)
	if err := this.Db.First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (this *User) Create(user *model.User) error {
	if err := this.Db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (this *User) Update(user *model.User) (*model.User, error) {
	if err := this.Db.Model(user).Update(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (this *User) FindByField(key string, value string, fields string) (*model.User, error) {
	if len(fields) == 0 {
		fields = "*"
	}
	user := &model.User{}
	if err := this.Db.Select(fields).Where(key+" = ?", value).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
