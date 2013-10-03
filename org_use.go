package main

import (
	"log"
)

// space.use also sets org
func init_deprecated() {
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

	i, err := choose(OrgList(orgs))
	if err != nil {
		log.Fatal(err)
	}

	target.Org = orgs[i].Name
	target.OrgGuid = orgs[i].Guid

	return
}
