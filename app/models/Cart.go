package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/jinzhu/gorm"
)

type Cart struct {
	gorm.Model
	CartProducts []CartProduct
	UserID       uint `schema:"user_id"`
}

/**
*
* Create cart
**/
func (cart *Cart) Create() (err error) {
	err = database.SQL.Create(&cart).Error
	return
}

/**
*
* Total price cart
**/
func (cart *Cart) TotalPrice() uint {
	sum := 0
	for _, v := range cart.CartProducts {
		sum += int(v.Quantity) * v.Product.Price
	}
	return uint(sum + 20000)
}

/**
*
* Delete
**/
func (cart *Cart) Delete() error {
	err := database.SQL.Delete(&cart).Error
	return err
}

/**
*
* Find cart by user id
**/
func (cart *Cart) FindByUserID(userID uint) error {
	var err error
	err = database.SQL.Where("user_id = ?", userID).First(&cart).Error
	return err
}
