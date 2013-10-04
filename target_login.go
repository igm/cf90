package main

import (
	"fmt"
	"github.com/igm/cf90/echo"
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
	alias := params["alias"]

	target, err := c.GetTargetByAlias(alias)
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
	fmt.Printf("Password: ")
	echo.EchoOff(func() {
		fmt.Scanf("%s\n", &password)
	})
	// password, err = gopass.GetPass("Password: ")

	err = target.Login(login, password)
	if err != nil {
		log.Fatal(err)
	}
	return
}
