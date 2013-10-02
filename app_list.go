package main

import (
	"fmt"
	"log"
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
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	summary, err := target.Summary(c.data.ActiveSpace)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name                     Status      Usage      URL")
	for _, app := range summary.Apps {
		instances := fmt.Sprintf("%d x %dM", app.Instances, app.Memory)
		fmt.Printf("%-25s%-12s%-11s%-13s\n", app.Name, app.State, instances, app.Urls)
	}
}
