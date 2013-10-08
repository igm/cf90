package main

import (
	"fmt"
	"github.com/igm/cf"
	"log"
	"time"
)

func init() {
	register(&Command{
		group:  "Application",
		name:   "app.list",
		help:   "Show a list of apps",
		handle: app_list,
	})
}

func app_list() {
	ch := make(chan []cf.App)
	go func() {
		target, err := c.SelectedTarget()
		if err != nil {
			log.Fatal(err)
		}
		apps, err := target.AppsGet()
		if err != nil {
			log.Fatal(err)
		}
		ch <- apps
	}()

	go func() {
		<-time.After(20 * time.Second)
		log.Fatal("request timeout!")
	}()

	fmt.Print("waiting for response")
	for {
		fmt.Print(".")
		select {
		case apps := <-ch:
			fmt.Println("done")
			list(AppList(apps))
			goto end
		case <-time.After(500 * time.Millisecond):
		}
	}
end:
}
