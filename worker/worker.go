package worker


import (

	"fmt"
	"github.com/robfig/cron"
	"time"
)

//v0.0.1
func registerWorker(){
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

	//todo implement fetching data
	//todo implement saving fetched data


}