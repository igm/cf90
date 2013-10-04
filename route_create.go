package main

import (
	"fmt"
	"log"
)

func init() {
	register(&Command{
		group: "Route",
		name:  "route.create",
		help:  "Create a route in current space",
		params: []Param{
			Param{name: "host", desc: "Host name"},
			Param{name: "domain", desc: "Domain name"},
			Param{name: "space", desc: "Space name"},
			Param{name: "org", desc: "Organization name"},
		},
		handle: route_create,
	})
}

func route_create() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}

	domains, err := target.DomainsGet()
	if err != nil {
		log.Fatal(err)
	}
	i, err := DomainList(domains).FindOrPick(params["domain"])
	if err != nil {
		log.Fatal(err)
	}
	domain := domains[i]

	spaces := domain.Spaces
	i, err = SpaceList(spaces).FindOrPick(params["space"], params["org"])
	if err != nil {
		log.Fatal(err)
	}
	space := spaces[i]

	name, exists := params["host"]
	if !exists {
		fmt.Print("Host name: ")
		_, err = fmt.Scanf("%s\n", &name)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = target.RouteCreate(name, domain.Guid, space.Guid)
	if err != nil {
		log.Fatal(err)
	}
}
