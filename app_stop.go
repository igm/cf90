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
	apps, err := target.AppsGet()
	if err != nil {
		log.Fatal(err)
	}

	name, appId := params["name"], ""
	for _, app := range apps {
		if app.Name == name {
			appId = app.Guid
		}
	}

	if appId == "" {
		i, err := choose(AppList(apps))
		if err != nil {
			log.Fatal(err)
		}
		appId = apps[i].Guid
	}

	err = target.AppStop(appId)
	if err != nil {
		log.Fatal(err)
	}

}
