package api

import (
	fetchers "Kamil-Ambroziak"
	"Kamil-Ambroziak/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Api struct{
	Storage fetchers.Storage
	Worker fetchers.Worker
}

func NewAPIServer(mySqlClient fetchers.Storage, worker fetchers.Worker) {
	api := &Api{
		Storage: mySqlClient,
		Worker: worker,
	}
	router := gin.Default()
	router.Use(checkSize)
	router.POST("/api/fetcher", api.AddFetcher)
	router.PATCH("/api/fetcher/:id", api.UpdateFetcher)
	router.DELETE("/api/fetcher/:id", api.DeleteFetcher)
	router.GET("/api/fetcher", api.GetAllFetchers)
	router.GET("/api/fetcher/:id/history", api.GetHistoryForFetcher)
	router.Run(":8080")
}

func checkSize(c *gin.Context){
	size := c.Request.ContentLength
	if size > 10 {
		c.AbortWithStatusJSON(http.StatusRequestEntityTooLarge,utils.NewRestError("entity too large",http.StatusRequestEntityTooLarge,"payload can be max 1024 bytes",nil))
		return
	}
	c.Next()
}
