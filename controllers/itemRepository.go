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

type ItemRepo struct {
	Db *gorm.DB
}

func NewRepoItem() *ItemRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Item{})
	return &ItemRepo{Db: db}
}

// create user
func (repository *ItemRepo) CreateItem(c *gin.Context) {
	var item models.Item
	c.BindJSON(&item)
	err := models.CreateItem(repository.Db, &item)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, item)
}

// get users
func (repository *ItemRepo) GetItems(c *gin.Context) {
	var item []models.Item
	err := models.GetItems(repository.Db, &item)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, item)
}

// get users
func (repository *OrderRepo) GetItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var item models.Item
	err := models.GetItem(repository.Db, &item, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, item)
}
