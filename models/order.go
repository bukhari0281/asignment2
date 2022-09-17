package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func GetOrders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Ari",
		"bio":  "Tukang nulis",
	})
}

func GetOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Ari",
		"bio":  "Tukang nulis",
	})
}

func CreateOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Ari",
		"bio":  "Tukang nulis",
	})
}

func UpdateOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Ari",
		"bio":  "Tukang nulis",
	})
}

func DeleteOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Ari",
		"bio":  "Tukang nulis",
	})
}
