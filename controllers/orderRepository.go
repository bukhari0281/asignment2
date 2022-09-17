package controllers

import (
	"errors"
	"net/http"
	"one-million-data/database"
	"one-million-data/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderRepo struct {
	Db *gorm.DB
}

func NewRepo() *OrderRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Order{})
	return &OrderRepo{Db: db}
}

// create user
func (repository *OrderRepo) CreateOrder(c *gin.Context) {
	var order models.Order
	c.BindJSON(&order)
	err := models.CreateOrder(repository.Db, &order)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, order)
}

// get users
func (repository *OrderRepo) GetOrders(c *gin.Context) {
	var order []models.Order
	err := models.GetOrders(repository.Db, &order)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, order)
}

// get users
func (repository *OrderRepo) GetOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var order models.Order
	err := models.GetOrder(repository.Db, &order, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, order)
}

// get users
func (repository *OrderRepo) UpdateOrder(c *gin.Context) {
	var order models.Order
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetOrder(repository.Db, &order, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&order)
	err = models.UpdateOrder(repository.Db, &order)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, order)
}

// get users
func (repository *OrderRepo) DeleteOrder(c *gin.Context) {
	var order models.Order
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteOrder(repository.Db, &order, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, order)
}
