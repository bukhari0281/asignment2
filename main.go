package main

import (
	"one-million-data/models"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
//
//	func getAlbums(c *gin.Context) {
//		c.IndentedJSON(http.StatusOK, albums)
//	}
type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var books = []book{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	// v1 := r.Group("/v1")

	r.GET("/", models.GetOrders)
	r.GET("/order/:id", models.GetOrder)
	r.GET("/order", models.CreateOrder)
	r.GET("/order/:id", models.UpdateOrder)
	r.GET("/order/;id", models.DeleteOrder)
	r.Run()

}
