package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Item_Id     uint `gorm:"primaryKey"`
	Item_Code   uint
	Description string
	Quantity    uint
	Order_Id    uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func GetItems(db *gorm.DB, Item *[]Item) (err error) {
	err = db.Find(Item).Error
	if err != nil {
		return err
	}
	return nil
}

func GetItem(db *gorm.DB, Item *Item, id int) (err error) {
	err = db.Where("id = ?").First(Item).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateItem(db *gorm.DB, Item *Item) (err error) {
	err = db.Create(Item).Error
	if err != nil {
		return err
	}
	return nil
}
