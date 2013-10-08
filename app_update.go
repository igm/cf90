package main

import (
	"fmt"
	"github.com/igm/cf"
	"log"
	"strconv"
)

func init() {
	register(&Command{
		group: "Application",
		name:  "app.update",
		help:  "Update application",
		params: []Param{
			Param{name: "name", desc: "Application name"},
			Param{name: "space", desc: "Space name"},
			Param{name: "org", desc: "Organization name"},
			Param{name: "mem", desc: "Memory allocation [MB]"},
			Param{name: "instances", desc: "Number of instances"},
		},
		handle: app_update,
	})
}

func app_update() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	// Get Application ID
	app, err := target.AppFind(params["name"], params["space"], params["org"])
	if err != nil {
		log.Fatal(err)
	}
	update := &cf.AppUpdate{
		Guid:      app.Guid,
		Instances: app.Instances,
		Memory:    app.Memory,
	}

	instances, err := strconv.ParseInt(params["instances"], 0, 32)
	if err != nil {
		fmt.Printf("Number of instances [current %d]: ", update.Instances)
		_, err = fmt.Scanf("%d\n", &update.Instances)
		if err != nil {
			log.Println("not changed")
		}
	} else {
		update.Instances = int(instances)
	}

	memory, err := strconv.ParseInt(params["mem"], 0, 32)
	if err != nil {
		fmt.Printf("Memory in MB [current %d]: ", update.Memory)
		_, err = fmt.Scanf("%d\n", &update.Memory)
		if err != nil {
			log.Println("not changed")
		}
	} else {
		update.Memory = int(memory)
	}

	appUpdated, err := target.AppUpdate(update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(appUpdated)
}
