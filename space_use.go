package main

import (
	"log"
)

func init() {
	register(&Command{
		group:  "Space",
		name:   "space.use",
		help:   "Set default space",
		params: []Param{Param{name: "space", desc: "Space name"}},
		handle: space_use,
	})
}

func space_use() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}

	spaces, err := target.SpacesGet(target.OrgGuid)
	if err != nil {
		log.Fatal(err)
	}
	i, err := choose(SpaceList(spaces))
	if err != nil {
		log.Fatal(err)
	}
	target.Space = spaces[i].Name
	target.SpaceGuid = spaces[i].Guid
}
