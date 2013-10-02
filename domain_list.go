package main

import (
	"fmt"
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

	domains, err := target.DomainsGet(c.data.ActiveSpace)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Domains:")
	for _, domain := range domains {
		fmt.Printf("  %s\n", domain.Name)
	}
}
