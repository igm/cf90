package main

import (
	"fmt"
)

func init() {
	register(&Command{
		group:  "Target",
		name:   "target.list",
		help:   "Show known targets",
		handle: target_list,
	})
}

func target_list() {
	list(TargetList(c.data.Targets))
	if selected, err := c.SelectedTarget(); err == nil {
		fmt.Println("\nCurrent target:", selected.TargetUrl)
	}
}
