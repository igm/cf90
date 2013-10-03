package main

import (
	"log"
)

func init() {
	register(&Command{
		group:  "Service",
		name:   "service.list",
		help:   "Show a list of services",
		handle: service_list,
	})
}

func service_list() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	services, err := target.GetServices()
	if err != nil {
		log.Fatal(err)
	}
	list(ServiceList(services))
}
