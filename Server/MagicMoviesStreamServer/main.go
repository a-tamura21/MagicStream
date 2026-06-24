package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	controller "github.com/a-tamura21/MagicStream/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Welcome to streaming")
	})

	router.GET("/movies", controller.GetMovies())
	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())
	router.POST("/register", controller.RegisterUser())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Server failed to start", err)
	}
}
