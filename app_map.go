package main

import (
	"log"
)

func init() {
	register(&Command{
		group: "Application",
		name:  "app.map",
		help:  "Map application to given host and domain (route must already exist)",
		params: []Param{
			Param{name: "name", desc: "Application name"},
			Param{name: "host", desc: "Host name"},
			Param{name: "domain", desc: "Domain name"},
		},
		handle: app_map,
	})
}

func app_map() {

	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	summary, err := target.Summary(target.SpaceGuid)
	if err != nil {
		log.Fatal(err)
	}

	routes, err := target.RoutesGet()
	if err != nil {
		log.Fatal(err)
	}

	name, appId := params["name"], ""
	for _, app := range summary.Apps {
		if app.Name == name {
			appId = app.Guid
		}
	}

	if appId == "" {
		i, err := choose(AppList(summary.Apps))
		if err != nil {
			log.Fatal(err)
		}
		appId = summary.Apps[i].Guid
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
		index, err := choose(RouteList(routes))
		if err != nil {
			log.Fatal(err)
		}
		routeGUID = routes[index].Guid
	}

	err = target.AppAddRoute(appId, routeGUID)
	if err != nil {
		log.Fatal(err)
	}

}
