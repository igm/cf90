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

	routes, err := target.RoutesGet()
	if err != nil {
		log.Fatal(err)
	}

	host := params["host"]
	domain := params["domain"]

	routeGUID := ""
	for _, route := range routes {
		if route.Host == host && route.Domain.Name == domain {
			routeGUID = route.Guid
		}
	}

	if routeGUID == "" {
		// ADD console UI to pick host/domain
		log.Fatal("route not found")
	}

	err = target.RouteDelete(routeGUID)
	if err != nil {
		log.Fatal(err)
	}
}
