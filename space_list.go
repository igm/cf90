package main

import (
	"log"
)

func init() {
	register(&Command{
		group:  "Space",
		name:   "space.list",
		help:   "Show all spaces in organization",
		handle: space_list,
	})
}

func space_list() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}

	spaces, err := target.SpacesGet(target.OrgGuid)
	if err != nil {
		log.Fatal(err)
	}
	list(SpaceList(spaces))
}
