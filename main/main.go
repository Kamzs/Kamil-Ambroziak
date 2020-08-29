package main

import (
	"Kamil-Ambroziak/api"
	"Kamil-Ambroziak/logger"
	"Kamil-Ambroziak/storage"
	worker2 "Kamil-Ambroziak/worker"
	"github.com/pkg/errors"
)


func main() {

	msql, err := storage.NewMySQL()
	if err != nil {
		err = errors.Wrap(err, "main: ")
		log := logger.GetLogger()
		log.Print(err)
		return
	}

	worker := worker2.NewWorker()

	api.NewAPIServer(msql,worker)



}
