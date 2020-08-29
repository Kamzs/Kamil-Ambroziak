package main

import (
	"Kamil-Ambroziak/api"
	"Kamil-Ambroziak/logger"
	"Kamil-Ambroziak/storage"
	"Kamil-Ambroziak/worker"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	router = gin.Default()
)
func main() {

	msql, err := storage.NewMySQL()
	if err != nil {
		err = errors.Wrap(err, "main: ")
		log := logger.GetLogger()
		log.Print(err)
		return
	}

	worker := worker.NewWorker()
	api:= api.NewAPIServer(msql,worker)

	router.POST("/api/fetcher", api.AddFetcher)
	router.PATCH("/api/fetcher/:id", api.UpdateFetcher)
	router.DELETE("/api/fetcher/:id", api.DeleteFetcher)
	router.GET("/api/fetcher", api.GetAllFetchers)

	router.GET("/api/fetcher/:id/history", api.GetHistoryForFetcher)
	//api.NewAPIServer(mySqlClient)
	router.Run(":8080")

}
