package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func AddFetcher(c *gin.Context) {

	//todo middleware do contet lenghta
	checkLength(c.Request.ContentLength)


	//todo add timeouts (przez context??)
	context.Background()
	client:= http.Client{Timeout: 300}
	resp, err:= client.Get("url")
	htmlData, err := ioutil.ReadAll(resp.Body)
	data:= string(htmlData)

	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	user, err := services.UsersService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}
