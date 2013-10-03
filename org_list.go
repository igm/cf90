package main

import (
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

	list(OrgList(orgs))
}
