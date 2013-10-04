package main

import (
	"fmt"
	"log"
	"time"
)

func init() {
	register(&Command{
		group: "Application",
		name:  "app.info",
		help:  "Show application info",
		params: []Param{
			Param{name: "name", desc: "Application name"},
		},
		handle: app_info,
	})
}

func app_info() {
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

	instances, err := target.AppInstances(appId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%-6s %-7s %-25s %-16s %-6s %-16s %-6s\n", "Inst#", "State", "Since", "Debug IP", "Debug Port", "Console IP", "Console Port")
	for instance, detail := range instances {
		date := time.Unix(int64(detail.Since), 0)
		dates := date.Format(time.RFC3339)
		fmt.Printf("(%-2s)   %-7s %-25s %-16s %-6s %-16s %-6s\n", instance, detail.State, dates, detail.DebugIP, detail.DebugPort, detail.ConsoleIP, detail.ConsolePort)
	}
}
