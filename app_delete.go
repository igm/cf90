package main

import (
	"log"
)

func init() {
	register(&Command{
		group:  "Application",
		name:   "app.delete",
		help:   "Delete application",
		params: []Param{Param{name: "name", desc: "Application name"}},
		handle: app_delete,
	})
}

func app_delete() {

	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	summary, err := target.Summary(target.SpaceGuid)
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
		i, err := choose(AppList(summary.Apps))
		if err != nil {
			log.Fatal(err)
		}
		appId = summary.Apps[i].Guid
	}

	err = target.AppDelete(appId)
	if err != nil {
		log.Fatal(err)
	}
}
