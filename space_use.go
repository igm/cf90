package main

import (
	"fmt"
	"github.com/igm/cf"
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

	orgGUID, err := c.SelectedOrg()
	if err != nil {
		log.Fatal(err)
	}

	spaces, err := target.SpacesGet(orgGUID)
	if err != nil {
		log.Fatal(err)
	}

	space, err := chooseSpace(spaces)
	if err != nil {
		log.Fatal(err)
	}
	c.data.ActiveSpace = space.Guid
}

func chooseSpace(spaces []cf.Space) (space cf.Space, err error) {
	fmt.Println("Spaces:")
	for i, space := range spaces {
		active := ""
		if c.data.ActiveSpace == space.Guid {
			active = "[current]"
		}
		fmt.Printf("  (%d) %s %s\n", i+1, space.Name, active)
	}

	fmt.Print("Select space: ")
	var sp int
	if _, err = fmt.Scanf("%d ", &sp); err == nil {
		space = spaces[sp-1]
		// c.data.ActiveSpace = spaces[sp-1].Guid
	}
	return
}
