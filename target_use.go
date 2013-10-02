package main

import (
	"code.google.com/p/gopass"
	"fmt"
	"github.com/igm/cf"
	"log"
)

func init() {
	register(&Command{
		group:  "Target",
		name:   "target.use",
		help:   "Set current target",
		params: []Param{Param{name: "target", desc: "Target URL"}},
		handle: target_use,
	})
}

func target_use() {
	targetUrl, exists := params["target"]

	if !exists {
		fmt.Print("Target URL: ")
		fmt.Scanf("%s\n", &targetUrl)
	}

	if c.Select(targetUrl) != nil {
		var login, password string
		fmt.Printf("Login: ")
		fmt.Scanf("%s\n", &login)
		password, err := gopass.GetPass("Password: ")
		target, err := cf.Login(targetUrl, login, password)
		if err != nil {
			log.Fatal(err)
		}
		c.AddTarget(target)
		c.Select(targetUrl)

	}
	org_use()
	space_use()

	return
}
