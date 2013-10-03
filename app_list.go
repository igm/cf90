package main

import (
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
	summary, err := target.Summary(target.SpaceGuid)
	if err != nil {
		log.Fatal(err)
	}
	list(AppList(summary.Apps))
}
