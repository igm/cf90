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

	err = target.AppStop(appId)
	if err != nil {
		log.Fatal(err)
	}

}
