package main

import (
	"io"
	"log"
	"os"
)

func init() {
	register(&Command{
		group: "Application",
		name:  "app.cat",
		help:  "Get application instance file content",
		params: []Param{
			Param{name: "name", desc: "Application name"},
			Param{name: "space", desc: "Space name"},
			Param{name: "org", desc: "Organization name"},
			Param{name: "instance", desc: "Application instance", defval: "0"},
			Param{name: "file", desc: "Remote file"},
		},
		handle: app_cat,
	})
}

func app_cat() {

	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	// Get Application ID
	app, err := target.AppFind(params["name"], params["space"], params["org"])
	if err != nil {
		log.Fatal(err)
	}
	reader, err := target.AppGet(app.Guid, params["instance"], params["file"])
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, reader)
}
