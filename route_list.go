package main

import (
	"fmt"
	"log"
)

func init() {
	register(&Command{
		group:  "Route",
		name:   "route.list",
		help:   "Show all routes",
		handle: route_list,
	})
}

func route_list() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}

	routes, err := target.RoutesGet()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Routes:")
	fmt.Printf("  %-15s %s\n", "Host", "Domain")
	for _, route := range routes {
		fmt.Printf("  %-15s %s\n", route.Host, route.Domain.Name)
	}
}
