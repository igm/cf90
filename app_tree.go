package main

import (
	"fmt"
	"log"
)

func init() {
	register(&Command{
		group: "Application",
		name:  "app.tree",
		help:  "Show application instance directory tree.",
		params: []Param{
			Param{name: "name", desc: "Application name"},
			Param{name: "org", desc: "Organization name"},
			Param{name: "space", desc: "Space name"},
			Param{name: "instance", desc: "Application instance", defval: "0"},
			Param{name: "dir", desc: "Remote dir"},
		},
		handle: app_tree,
	})
}

func app_tree() {

	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	// Get Application ID
	app, err := target.AppFind(params["name"], params["space"], params["org"])
	if err != nil {
		log.Fatal(err)
	}

	showTree(target, app.Guid, params["instance"], params["dir"], "")

}

func showTree(target *Target, appid, instance, dir, prefix string) {
	files, err := target.AppLs(appid, instance, dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.Dir {
			fmt.Printf("%s%s\n", prefix, file.Name)
			showTree(target, appid, instance, dir+file.Name, prefix+"  ")
		} else {
			fmt.Printf("%s%-7s %s\n", prefix, file.Size, file.Name)
		}
	}
}
