package main

import (
	"fmt"
	"log"
)

func init() {
	register(&Command{
		group:  "Organization",
		name:   "org.use",
		help:   "Set default organization",
		params: []Param{Param{name: "org", desc: "Organization name"}},
		handle: org_use,
	})
}

func org_use() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}

	orgs, err := target.OrganizationsGet()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Organizations:")
	for i, org := range orgs {
		active := ""
		if c.data.ActiveOrg == org.Guid {
			active = "[current]"
		}
		fmt.Printf("  (%d) %s %s\n", i+1, org.Name, active)
	}
	fmt.Print("Select organization: ")
	var org int
	if _, err = fmt.Scanf("%d ", &org); err == nil {
		c.data.ActiveOrg = orgs[org-1].Guid
	} else {
		log.Fatal(err)
	}
}
