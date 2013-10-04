package main

import (
	"log"
)

func init() {
	register(&Command{
		group: "Application",
		name:  "app.map",
		help:  "Map application to given host and domain (route must already exist)",
		params: []Param{
			Param{name: "name", desc: "Application name"},
			Param{name: "space", desc: "Space name"},
			Param{name: "org", desc: "Organization name"},

			Param{name: "host", desc: "Host name"},
			Param{name: "domain", desc: "Domain name"},
		},
		handle: app_map,
	})
}

func app_map() {

	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}

	// Get Application ID
	app, err := target.AppFind(params["name"], params["space"], params["org"])
	if err != nil {
		log.Fatal(err)
	}

	// Get RouteID
	route, err := target.RouteFind(params["host"], params["domain"])
	if err != nil {
		log.Fatal(err)
	}

	// Perform action
	err = target.AppAddRoute(app.Guid, route.Guid)
	if err != nil {
		log.Fatal(err)
	}

}
