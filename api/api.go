package api

import (
	fetchers "Kamil-Ambroziak"
)

type Api struct{
	Storage fetchers.Storage
}

func WithStorageClient(client fetchers.Storage) func(*Api) {
	return func(a *Api) {
		a.Storage = client
	}
}

func NewAPIServer(mySqlClient fetchers.Storage) *Api {
	return &Api{
		Storage: mySqlClient,
	}
}

