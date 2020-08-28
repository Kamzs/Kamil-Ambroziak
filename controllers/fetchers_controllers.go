package controllers

import (
	"Kamil-Ambroziak/domain/fetchers"
	"Kamil-Ambroziak/services"
	"Kamil-Ambroziak/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//todo add api struct which can be used to pass worker??

func getFetcherId(fetcherIdParam string) (int64, utils.RestErr) {
	fetcherId, fetcherErr := strconv.ParseInt(fetcherIdParam, 10, 64)
	if fetcherErr != nil {
		return 0, utils.NewBadRequestError("fetcher id should be an int64")
	}
	return fetcherId, nil
}

func AddFetcher(c *gin.Context) {
	var fetcher fetchers.Fetcher
	if err := c.ShouldBindJSON(&fetcher); err != nil {
		restErr := utils.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := services.FetchersService.CreateFetcher(fetcher)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	//todo add worker

	c.JSON(http.StatusCreated, JsonWithID{Id: result.Id})
}
func GetAllFetchers(c *gin.Context) {

	fetchers, getErr := services.FetchersService.FindAllFetchers()
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(http.StatusOK, fetchers)
}

//todo determine if pathing or full update
func UpdateFetcher(c *gin.Context) {
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

	result, err := services.FetchersService.UpdateFetcher(fetcher)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteFetcher(c *gin.Context) {
	fetcherId, idErr := getFetcherId(c.Param("id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := services.FetchersService.DeleteFetcher(fetcherId); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK,  JsonWithID{Id: fetcherId})
}

//todo for fetching history

func GetHistoryForFetcher(c *gin.Context) {

	userId, idErr := getFetcherId(c.Param("id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	fetchingHistory, getErr := services.FetchersService.GetHistoryForFetcher(userId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	//todo after implementation of new table for fetched data saving change output
	c.JSON(http.StatusOK, fetchingHistory)
}
