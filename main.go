package main

import (
	"Kamil-Ambroziak/controllers"
	"github.com/gin-gonic/gin"

)

var (
	router = gin.Default()
)

func main() {
	router.Run(":8080")
	mapUrls()
}
func mapUrls() {

	router.POST("/api/fetcher", controllers.AddFetcher)
	router.PATCH("/api/fetcher/:id", controllers.UpdateFetcher)
	router.DELETE("/api/fetcher/:id", controllers.DeleteFetcher)
	router.GET("/api/fetcher", controllers.GetAllFetchers)

	router.GET("/api/fetcher/:id/history", controllers.GetHistoryForFetcher)
}

