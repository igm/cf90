package main

import (
	"log"
)

func init() {
	register(&Command{
		group:  "Target",
		name:   "target.logout",
		help:   "Logout from target",
		params: []Param{Param{name: "target", desc: "Target URL"}},
		handle: target_logout,
	})
}

func target_logout() {

	target, err := c.GetTarget(params["target"])

	if err != nil {
		index, err := choose(TargetList(c.data.Targets))
		if err != nil {
			log.Fatal(err)
		}
		target = c.data.Targets[index]
	}
	target.Logout()
	return
}
