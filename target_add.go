package main

import (
	"log"
)

func init() {
	register(&Command{
		group: "Target",
		name:  "target.add",
		help:  "Add new target",
		params: []Param{
			Param{name: "target", desc: "Target URL"},
			Param{name: "alias", desc: "The alias for this target"},
		},
		handle: target_add,
	})
}

func target_add() {
	var err error
	targetUrl := params["target"]
	if targetUrl == "" {
		targetUrl, err = enterText("Enter target URL: ")
		if err != nil {
			log.Fatal(err)
		}
	}

	alias := params["alias"]
	if alias == "" {
		alias, err = enterText("Enter alias for this target: ")
		if err != nil {
			log.Fatal(err)
		}
	}

	c.AddTarget(NewTarget(targetUrl, alias))
}
