package api

import (
	fetchers "Kamil-Ambroziak"
	"github.com/gin-gonic/gin"
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
	router.POST("/api/fetcher", api.AddFetcher)
	router.PATCH("/api/fetcher/:id", api.UpdateFetcher)
	router.DELETE("/api/fetcher/:id", api.DeleteFetcher)
	router.GET("/api/fetcher", api.GetAllFetchers)
	router.GET("/api/fetcher/:id/history", api.GetHistoryForFetcher)
	router.Run(":8080")
}
