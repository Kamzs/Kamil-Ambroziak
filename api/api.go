package api

import (
	"net/http"

	fetchers "github.com/Kamzs/Kamil-Ambroziak"
	"github.com/Kamzs/Kamil-Ambroziak/utils"
	"github.com/gin-gonic/gin"
)

type Api struct {
	Storage fetchers.Storage
	Worker  fetchers.Worker
	Router  *gin.Engine
}

func NewAPIServer(mySqlClient fetchers.Storage, worker fetchers.Worker) *Api {
	api := &Api{
		Storage: mySqlClient,
		Worker:  worker,
		Router:  gin.Default(),
	}
	api.Router.Use(checkSize)
	api.Router.POST("/api/fetcher", api.AddFetcher)
	api.Router.PATCH("/api/fetcher/:id", api.UpdateFetcher)
	api.Router.DELETE("/api/fetcher/:id", api.DeleteFetcher)
	api.Router.GET("/api/fetcher", api.GetAllFetchers)
	api.Router.GET("/api/fetcher/:id/history", api.GetHistoryForFetcher)
	return api
}

func checkSize(c *gin.Context) {
	size := c.Request.ContentLength
	if size > 1024 {
		c.AbortWithStatusJSON(http.StatusRequestEntityTooLarge, utils.NewRestError("entity too large", http.StatusRequestEntityTooLarge, "payload can be max 1024 bytes", nil))
		return
	}
	c.Next()
}
