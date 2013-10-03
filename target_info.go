package main

import (
	"fmt"
	"log"
)

func init() {
	register(&Command{
		group:  "Target",
		name:   "target.info",
		help:   "Show target information",
		handle: target_info,
	})
}

func target_info() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	info, err := target.Info()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("")
	fmt.Println("Alias                  :", target.Alias)
	fmt.Println("URL                    :", target.TargetUrl)
	fmt.Println("Username               :", target.Username)
	fmt.Println("Organization           :", target.Org)
	fmt.Println("Space                  :", target.Space)
	fmt.Println("")
	fmt.Println("Name                   :", info.Name)
	fmt.Println("Build                  :", info.Build)
	fmt.Println("Support                :", info.Support)
	fmt.Println("Version                :", info.Version)
	fmt.Println("Description            :", info.Description)
	fmt.Println("Authorization Endpoint :", info.AuthorizationEndpoint)
	fmt.Println("API Version            :", info.ApiVersion)
	fmt.Println("")
}
