package main

import (
	"log"
)

func init() {
	register(&Command{
		group:  "Application",
		name:   "app.stop",
		help:   "Stop application",
		params: []Param{Param{name: "name", desc: "Application name"}},
		handle: app_stop,
	})
}

func app_stop() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	// Get Application ID
	app, err := target.AppFind(params["name"], params["space"], params["org"])
	if err != nil {
		log.Fatal(err)
	}
	_, err = target.AppStop(app.Guid)
	if err != nil {
		log.Fatal(err)
	}

}
