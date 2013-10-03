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
	apps, err := target.AppsGet()
	if err != nil {
		log.Fatal(err)
	}
	list(AppList(apps))
}
