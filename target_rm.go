package main

import (
	"log"
)

func init() {
	register(&Command{
		group:  "Target",
		name:   "target.rm",
		help:   "Remove target from the list of known targets",
		params: []Param{Param{name: "target", desc: "Target URL"}},
		handle: target_rm,
	})
}

func target_rm() {
	targetUrl := params["target"]

	target, err := c.GetTarget(targetUrl)
	if err != nil {
		i, err := choose(TargetList(c.data.Targets))
		if err != nil {
			log.Fatal(err)
		}
		target = c.data.Targets[i]
	}
	c.RemoveTarget(target)
}
