package api

import (
	fetchers "Kamil-Ambroziak"
	"Kamil-Ambroziak/utils"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
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

	//todo implement
	jobID, restErr := api.Worker.RegisterFetcher(&fetcher)
	if restErr != nil {
		c.JSON(restErr.Status(), restErr)
	}
	fetcher.JobID = int64(jobID)
	if err := api.Storage.SaveFetcher(&fetcher); err != nil {
		c.JSON(err.Status(), err)
		api.Worker.DeregisterFetcher(cron.EntryID(fetcher.JobID))
		return
	}

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

	var newFetcher fetchers.Fetcher
	if err := c.ShouldBindJSON(&newFetcher); err != nil {
		restErr := utils.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	oldFetcher, restErr := api.Storage.GetFetcher(fetcherId)
	if restErr !=nil {
		c.JSON(restErr.Status(), restErr)
	}

	oldJobId := oldFetcher.JobID

	fillMissingFields(oldFetcher,&newFetcher)

	newJobId, err := api.Worker.RegisterFetcher(&newFetcher)
	if err != nil {
		c.JSON(restErr.Status(), restErr)
	}
	api.Worker.DeregisterFetcher(cron.EntryID(oldJobId))

	newFetcher.Id = fetcherId
	newFetcher.JobID = int64(newJobId)
	err = api.Storage.UpdateFetcher(&newFetcher)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, JsonWithID{Id: newFetcher.Id})
}

func (api *Api) DeleteFetcher(c *gin.Context) {
	fetcherId, idErr := getFetcherId(c.Param("id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	fetcher, err := api.Storage.GetFetcher(fetcherId)
	if err != nil {
		c.JSON(err.Status(), err)
	}
	api.Worker.DeregisterFetcher(cron.EntryID(fetcher.JobID))

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
