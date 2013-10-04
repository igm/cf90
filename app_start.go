package main

import (
	"log"
)

func init() {
	register(&Command{
		group:  "Application",
		name:   "app.start",
		help:   "Start application",
		params: []Param{Param{name: "name", desc: "Application name"}},
		handle: app_start,
	})
}

func app_start() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	// Get Application ID
	app, err := target.AppFind(params["name"], params["space"], params["org"])
	if err != nil {
		log.Fatal(err)
	}

	err = target.AppStart(app.Guid)
	if err != nil {
		log.Fatal(err)
	}
}
