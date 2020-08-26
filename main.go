package main

import (
	"Kamil-Ambroziak/controllers"
	"fmt"
	"github.com/robfig/cron"
	"time"
	"github.com/gin-gonic/gin"

)

var (
	router = gin.Default()
)

func main() {

	router.Run(":8080")
	mapUrls()

	c := cron.New()

	c.AddFunc("*/1 * * * *", func() { fmt.Println("getPage1 1 sec") })

	// Start cron with one scheduled job
	c.Start()
	time.Sleep(2 * time.Minute)

	// Funcs may also be added to a running Cron
	//not working
	//	entryID2, _ := c.AddFunc("*/2 * * * *", func() { fmt.Println("getPage2 2 sec") })
	c.AddFunc("*/2 * * * *", func() { fmt.Println("getPage2 2 sec") })
	time.Sleep(5 * time.Minute)

	//Remove Job2 and add new Job2 that run every 1 minute
	//not working
	//c.Remove(entryID2)
	fmt.Println("getPage2 2 sec deleted")
	c.AddFunc("*/1 * * * *", func() { fmt.Println("getPage2 1 sec")})
	time.Sleep(5 * time.Minute)

}
func mapUrls(){
	router.POST("/api/fetcher", controllers.AddFetcher)

}

