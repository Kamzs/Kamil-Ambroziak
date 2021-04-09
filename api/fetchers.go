package api

import (
	"net/http"

	fetchers "github.com/Kamzs/Kamil-Ambroziak"
	"github.com/Kamzs/Kamil-Ambroziak/utils"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func (api *Api) AddFetcher(c *gin.Context) {

	var fetcher fetchers.Fetcher
	if err := c.ShouldBindJSON(&fetcher); err != nil {
		restErr := utils.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	if restErr := fetcher.Validate(false); restErr != nil {
		c.JSON(restErr.Status(), restErr)
		return
	}
	jobID, restErr := api.Worker.RegisterFetcher(&fetcher)
	if restErr != nil {
		c.JSON(restErr.Status(), restErr)
		return
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

	found, getErr := api.Storage.FindAllFetchers()
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	var resp []GetAllFetchersResponse
	for index, value := range found {
		resp = append(resp, GetAllFetchersResponse{})
		resp[index].Id = value.Id
		resp[index].Interval = value.Interval
		resp[index].Url = value.Url
	}
	c.JSON(http.StatusOK, resp)
}

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
	if restErr != nil {
		c.JSON(restErr.Status(), restErr)
		return
	}

	oldJobId := oldFetcher.JobID

	if restErr := newFetcher.Validate(true); restErr != nil {
		c.JSON(restErr.Status(), restErr)
		return
	}
	fillMissingFields(oldFetcher, &newFetcher)

	newJobId, restErr := api.Worker.RegisterFetcher(&newFetcher)
	if restErr != nil {
		c.JSON(restErr.Status(), restErr)
	}
	api.Worker.DeregisterFetcher(cron.EntryID(oldJobId))

	newFetcher.Id = fetcherId
	newFetcher.JobID = int64(newJobId)
	restErr = api.Storage.UpdateFetcher(&newFetcher)
	if restErr != nil {
		c.JSON(restErr.Status(), restErr)
		return
	}
	c.JSON(http.StatusOK, FetcherUpdateResponse{
		Id:       newFetcher.Id,
		Url:      newFetcher.Url,
		Interval: newFetcher.Interval,
	})
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
		return
	}
	api.Worker.DeregisterFetcher(cron.EntryID(fetcher.JobID))

	if err := api.Storage.DeleteFetcher(fetcherId); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, JsonWithID{Id: fetcherId})
}

func (api *Api) GetHistoryForFetcher(c *gin.Context) {

	fetcherId, idErr := getFetcherId(c.Param("id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	fetchersSlice, getErr := api.Storage.GetHistoryForFetcher(fetcherId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	var resp []HistoryElementResponse
	for index, value := range fetchersSlice {
		resp = append(resp, HistoryElementResponse{})
		resp[index].CreatedAt = value.CreatedAt
		resp[index].Duration = value.Duration
		resp[index].Response = value.Response
	}
	c.JSON(http.StatusOK, resp)
}
