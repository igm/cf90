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
			Param{name: "host", desc: "host name"},
			Param{name: "domain", desc: "domain name"},
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

	name, exists := params["host"]
	if !exists {
		fmt.Print("Host name: ")
		_, err = fmt.Scanf("%s\n", &name)
		if err != nil {
			log.Fatal(err)
		}
	}

	domainGUID := ""
	domainName := params["domain"]
	for _, domain := range domains {
		if domain.Name == domainName {
			domainGUID = domain.Guid
		}
	}

	if domainGUID == "" {
		index, err := choose(DomainList(domains))
		if err != nil {
			log.Fatal(err)
		}
		domainGUID = domains[index].Guid
	}

	err = target.RouteCreate(name, domainGUID, target.SpaceGuid)
	if err != nil {
		log.Fatal(err)
	}
}
