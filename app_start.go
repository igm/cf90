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
	summary, err := target.Summary(c.data.ActiveSpace)
	if err != nil {
		log.Fatal(err)
	}

	name, appId := params["name"], ""
	for _, app := range summary.Apps {
		if app.Name == name {
			appId = app.Guid
		}
	}

	if appId == "" {
		name, appId = chooseApplication(summary)
	}

	err = target.AppStart(appId)
	if err != nil {
		log.Fatal(err)
	}
}
