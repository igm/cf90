package main

import (
	"fmt"
	"github.com/igm/cf"
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
		routeGUID, err = chooseRoute(target)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = target.RouteDelete(routeGUID)
	if err != nil {
		log.Fatal(err)
	}
}

func chooseRoute(target *cf.Target) (routeGUID string, err error) {
	routes, err := target.RoutesGet()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Routes:")
	for i, route := range routes {
		fmt.Printf("  (%2d) %-15s %s\n", i+1, route.Host, route.Domain.Name)
	}
	fmt.Print("Select route: ")
	var index int
	if _, err = fmt.Scanf("%d\n", &index); err == nil && index > 0 && index <= len(routes) {
		routeGUID = routes[index-1].Guid
	}
	return
}
