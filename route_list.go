package main

import (
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

	list(RouteList(routes))
}
