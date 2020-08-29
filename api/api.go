package api

import (
	fetchers "Kamil-Ambroziak"
)

type Api struct{
	Storage fetchers.Storage
	Worker fetchers.Worker
}

func NewAPIServer(mySqlClient fetchers.Storage, worker fetchers.Worker) *Api {
	return &Api{
		Storage: mySqlClient,
		Worker: worker,
	}
}
