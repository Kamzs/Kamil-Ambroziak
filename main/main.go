package main

import (
	"github.com/Kamzs/Kamil-Ambroziak/api"
	"github.com/Kamzs/Kamil-Ambroziak/logger"
	"github.com/Kamzs/Kamil-Ambroziak/storage"
	"github.com/Kamzs/Kamil-Ambroziak/worker"
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
