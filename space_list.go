package main

import (
	"fmt"
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

	orgGUID, err := c.SelectedOrg()
	if err != nil {
		log.Fatal(err)
	}

	spaces, err := target.SpacesGet(orgGUID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Spaces:")
	for i, space := range spaces {
		active := ""
		if c.data.ActiveSpace == space.Guid {
			active = "[current]"
		}
		fmt.Printf("  (%d) %s %s\n", i+1, space.Name, active)
	}
}
