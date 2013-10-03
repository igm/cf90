package main

import ()

func init() {
	register(&Command{
		group: "Application",
		name:  "app.info",
		help:  "Show application info",
		params: []Param{
			Param{name: "name", desc: "Application name"},
		},
		handle: app_info,
	})
}

func app_info() {
	// TODO
}
