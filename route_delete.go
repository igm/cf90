package main

import (
	"log"
)

func init() {
	register(&Command{
		group: "Route",
		name:  "route.delete",
		help:  "Delete route",
		params: []Param{
			Param{name: "host", desc: "Host name"},
			Param{name: "domain", desc: "Domain name"},
		},
		handle: route_delete,
	})
}

func route_delete() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}

	route, err := target.RouteFind(params["host"], params["domain"])
	if err != nil {
		log.Fatal(err)
	}

	err = target.RouteDelete(route.Guid)
	if err != nil {
		log.Fatal(err)
	}
}
