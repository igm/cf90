package main

import (
	"fmt"
)

func init() {
	register(&Command{
		group:  "",
		name:   "version",
		help:   "show version information",
		handle: version,
	})
}

func version() {
	fmt.Println(versionInfo)
}

const versionInfo = `version: 0.1`
