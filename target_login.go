package main

import (
	"code.google.com/p/gopass"
	"fmt"
	"log"
)

func init() {
	register(&Command{
		group:  "Target",
		name:   "target.login",
		help:   "Login to target",
		params: []Param{Param{name: "alias", desc: "Target alias"}},
		handle: target_login,
	})
}

func target_login() {
	targetUrl := params["alias"]

	target, err := c.GetTarget(targetUrl)
	if err != nil {
		index, err := choose(TargetList(c.data.Targets))
		if err != nil {
			log.Fatal(err)
		}
		target = c.data.Targets[index]
	}

	var login, password string
	fmt.Printf("Login: ")
	fmt.Scanf("%s\n", &login)
	password, err = gopass.GetPass("Password: ")

	err = target.Login(login, password)
	if err != nil {
		log.Fatal(err)
	}
	return
}
