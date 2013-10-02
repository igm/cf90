package main

import (
	"fmt"
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

	fmt.Printf("\n%-15s %-9s %-15s %s\n", "Service", "Version", "Provider", "Description")
	for _, service := range services {
		fmt.Printf("%-15s %-9s %-15s %s\n", service.Label, service.Version, service.Provider, service.Description)
	}
}
