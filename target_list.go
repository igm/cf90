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
	fmt.Println("Known Targets:")
	selected, _ := c.SelectedTarget()
	for i, target := range c.data.Targets {
		if target == selected {
			fmt.Printf("  (%d) %s [current]\n", i+1, target.TargetUrl)
		} else {
			fmt.Printf("  (%d) %s \n", i+1, target.TargetUrl)
		}
	}
}
