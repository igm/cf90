package main

import (
	"log"
)

func init() {
	register(&Command{
		group:  "Domain",
		name:   "domain.list",
		help:   "Show a list of domains",
		handle: domain_list,
	})
}

func domain_list() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	domains, err := target.DomainsGet(target.SpaceGuid)
	if err != nil {
		log.Fatal(err)
	}
	list(DomainList(domains))
}
