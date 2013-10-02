package main

import (
	"fmt"
	"log"
)

func init() {
	register(&Command{
		group:  "Organization",
		name:   "org.list",
		help:   "Show all organizations",
		handle: org_list,
	})
}

func org_list() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}

	orgs, err := target.OrganizationsGet()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Organizations:")
	for _, org := range orgs {
		active := ""
		if c.data.ActiveOrg == org.Guid {
			active = "[current]"
		}
		fmt.Printf("  %s %s\n", org.Name, active)
	}

}
