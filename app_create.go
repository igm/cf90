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
		name:  "app.create",
		help:  "Create new application",
		params: []Param{
			Param{name: "name", desc: "Application name"},
			Param{name: "org", desc: "Organization name"},
			Param{name: "space", desc: "Space name"},
			Param{name: "buildpack", desc: "URL of the custom buildpack to use"},
			Param{name: "mem", desc: "Memory allocation [MB]"},
			Param{name: "instances", desc: "Number of instances"},
		},
		handle: app_create,
	})
}

func app_create() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}

	space, err := target.SpaceFind(params["space"], params["org"])
	if err != nil {
		log.Fatal(err)
	}

	appname, exists := params["name"]
	if !exists {
		fmt.Printf("Application name: ")
		fmt.Scanf("%s\n", &appname)
	}

	memory, err := strconv.ParseInt(params["mem"], 0, 32)
	if err != nil {
		fmt.Printf("Memory allocation in MB: ")
		_, err = fmt.Scanf("%d\n", &memory)
		if err != nil {
			log.Fatal(err)
		}
	}

	instances, err := strconv.ParseInt(params["instances"], 0, 32)
	if err != nil {
		fmt.Printf("Number of instances: ")
		_, err = fmt.Scanf("%d\n", &instances)
		if err != nil {
			log.Fatal(err)
		}
	}

	var bp *string = nil

	if params["buildpack"] != "" {
		bpack := params["buildpack"]
		bp = &bpack
	}

	_, err = target.AppCreate(&cf.NewApp{
		SpaceGUID: space.Guid,
		Name:      appname,
		Memory:    int(memory),
		Instances: int(instances),
		Buildpack: bp,
	})
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Application created: ", app.Name)
}
