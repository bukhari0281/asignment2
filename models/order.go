package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	id            uint `json:"id" gorm:"primary_key"`
	Customer_Name string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
	Items         []Item `gorm:"foreignKey:id"`
}

func GetOrders(db *gorm.DB, Order *[]Order) (err error) {
	err = db.Find(Order).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOrder(db *gorm.DB, Order *Order, id int) (err error) {
	err = db.Where("id = ?").First(Order).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateOrder(db *gorm.DB, Order *Order, Item *Item) (err error) {
	err = db.Create(Order).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateOrder(db *gorm.DB, Order *Order) (err error) {
	db.Save(Order)
	return nil
}

func DeleteOrder(db *gorm.DB, Order *Order, id int) (err error) {
	db.Where("id = ?").Delete(Order)
	return nil
}
