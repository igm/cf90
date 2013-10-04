package main

import (
	"log"
)

func init() {
	register(&Command{
		group: "Application",
		name:  "app.delete",
		help:  "Delete application",
		params: []Param{
			Param{name: "name", desc: "Application name"},
			Param{name: "space", desc: "Space name"},
			Param{name: "org", desc: "Organization name"},
		},
		handle: app_delete,
	})
}

func app_delete() {

	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	// Get Application ID
	app, err := target.AppFind(params["name"], params["space"], params["org"])
	if err != nil {
		log.Fatal(err)
	}

	err = target.AppDelete(app.Guid)
	if err != nil {
		log.Fatal(err)
	}
}
