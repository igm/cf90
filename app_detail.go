package main

import (
	"fmt"
	"log"
	"time"
)

func init() {
	register(&Command{
		group: "Application",
		name:  "app.detail",
		help:  "Show application info",
		params: []Param{
			Param{name: "name", desc: "Application name"},
			Param{name: "org", desc: "Organization name"},
			Param{name: "space", desc: "Spce name"},
		},
		handle: app_detail,
	})
}

func app_detail() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	// Get Application ID
	app, err := target.AppFind(params["name"], params["space"], params["org"])
	if err != nil {
		log.Fatal(err)
	}

	instances, err := target.AppInstances(app.Guid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%-6s %-9s %-25s %-16s %-6s %-16s %-6s\n", "Inst#", "State", "Since", "Debug IP", "Debug Port", "Console IP", "Console Port")
	for instance, detail := range instances {
		date := time.Unix(int64(detail.Since), 0)
		dates := date.Format(time.RFC3339)
		fmt.Printf("(%-2s)   %-9s %-25s %-16s %-6s %-16s %-6s\n", instance, detail.State, dates, detail.DebugIP, detail.DebugPort, detail.ConsoleIP, detail.ConsolePort)
	}
}
