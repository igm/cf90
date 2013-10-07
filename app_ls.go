package main

import (
	"fmt"
	"log"
)

func init() {
	register(&Command{
		group: "Application",
		name:  "app.ls",
		help:  "List application instance directory content.",
		params: []Param{
			Param{name: "name", desc: "Application name"},
			Param{name: "space", desc: "Space name"},
			Param{name: "org", desc: "Organization name"},
			Param{name: "instance", desc: "Application instance", defval: "0"},
			Param{name: "dir", desc: "Remote directory"},
		},
		handle: app_ls,
	})
}

func app_ls() {

	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	// Get Application ID
	app, err := target.AppFind(params["name"], params["space"], params["org"])
	if err != nil {
		log.Fatal(err)
	}
	files, err := target.AppLs(app.Guid, "0", params["dir"])
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Printf("%-10s %s\n", file.Size, file.Name)
	}
	// fmt.Print(files)
}

/*
GET /v2/apps/baf5421e-6b6d-479d-b3b4-7fa7f0ea37e1/instances/0/files/app/ HTTP/1.1
Host: api.run.pivotal.io
Accept: application/json
Authorization: *
User-Agent: cf90



HTTP/1.1 200 OK
Content-Length: 376
Connection: keep-alive
Content-Type: text/plain
Date: Mon, 07 Oct 2013 09:03:12 GMT
*/
