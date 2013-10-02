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

	spaceGUID := c.data.ActiveSpace
	if spaceGUID == "" {
		log.Fatal("no spaces selected.")
	}
	domains, err := target.DomainsGet(spaceGUID)
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
		fmt.Println("Domains:")
		for i, domain := range domains {
			fmt.Printf(" (%d) %s\n", i+1, domain.Name)
		}
		fmt.Print("Select domain: ")
		var selection int
		if _, err = fmt.Scanf("%d\n", &selection); err != nil {
			log.Fatal(err)
		} else {
			domainGUID = domains[selection-1].Guid
		}
	}

	err = target.RouteCreate(name, domainGUID, spaceGUID)
	if err != nil {
		log.Fatal(err)
	}
}
