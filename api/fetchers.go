package api

import (
	fetchers "Kamil-Ambroziak"
	"Kamil-Ambroziak/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (api *Api) AddFetcher(c *gin.Context) {
	var fetcher fetchers.Fetcher
	if err := c.ShouldBindJSON(&fetcher); err != nil {
		restErr := utils.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	if err := fetcher.Validate(); err != nil {
		restErr := utils.NewBadRequestError("validation failed")
		c.JSON(restErr.Status(), restErr)
		return
	}
	if err := api.Storage.SaveFetcher(&fetcher); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	//todo implement
	//registerWorker()

	//todo add worker

	c.JSON(http.StatusCreated, JsonWithID{Id: fetcher.Id})
}
func (api *Api) GetAllFetchers(c *gin.Context) {

	fetchers, getErr := api.Storage.FindAllFetchers()
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, fetchers)
}

//todo determine if pathing or full update
func (api *Api) UpdateFetcher(c *gin.Context) {
	fetcherId, idErr := getFetcherId(c.Param("id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var fetcher fetchers.Fetcher
	if err := c.ShouldBindJSON(&fetcher); err != nil {
		restErr := utils.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	fetcher.Id = fetcherId
	//todo should return updated row
	err := api.Storage.UpdateFetcher(&fetcher)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	//todo should return updated row
	c.JSON(http.StatusOK, nil)
}

func (api *Api) DeleteFetcher(c *gin.Context) {
	fetcherId, idErr := getFetcherId(c.Param("id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := api.Storage.DeleteFetcher(fetcherId); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, JsonWithID{Id: fetcherId})
}

//todo for fetching history

func (api *Api) GetHistoryForFetcher(c *gin.Context) {

	fetcherId, idErr := getFetcherId(c.Param("id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	fetcher, getErr := api.Storage.GetHistoryForFetcher(fetcherId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	//todo after implementation of new table for fetched data saving change output
	c.JSON(http.StatusOK, fetcher)
}
