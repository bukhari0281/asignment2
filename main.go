package main

import (
	"net/http"
	"one-million-data/controllers"
	"one-million-data/database"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
//
//	func getAlbums(c *gin.Context) {
//		c.IndentedJSON(http.StatusOK, albums)
//	}

func main() {
	r := setupRouter()

	database.ConnectDatabase()

	_ = r.Run(":8080")

}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	orderRepo := controllers.NewRepo()
	itemRepo := controllers.NewRepoItem()

	r.GET("/", orderRepo.GetOrders)
	r.POST("/order", orderRepo.CreateOrder)
	r.POST("/item", itemRepo.CreateItem)
	r.GET("/order/:id", orderRepo.GetOrder)
	r.PUT("/order/id", orderRepo.UpdateOrder)
	r.DELETE("/order/id", orderRepo.DeleteOrder)

	return r
}
