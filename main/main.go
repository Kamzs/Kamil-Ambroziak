package main

import (
	"Kamil-Ambroziak/api"
	"Kamil-Ambroziak/logger"
	"Kamil-Ambroziak/storage"
	"Kamil-Ambroziak/worker"
)


func main() {

	msql, err := storage.NewMySQL()
	if err != nil {
		log := logger.GetLogger()
		log.Print(err)
		return
	}
	w := worker.NewWorker()
	a := api.NewAPIServer(msql, w)
	a.Router.Run(":8080")

}
