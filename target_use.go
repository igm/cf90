package main

import (
	"errors"
	"fmt"
	"log"
)

func init() {
	register(&Command{
		group:  "Target",
		name:   "target.use",
		help:   "Set current target",
		params: []Param{Param{name: "alias", desc: "Target alias"}},
		handle: target_use,
	})
}

func target_use() {
	alias := params["alias"]

	idx, err := c.TargetByAlias(alias)
	if err != nil {
		log.Println(err)
		idx, err = choose(TargetList(c.data.Targets))
		if err != nil {
			log.Fatal(err)
		}
	}

	c.data.ActiveTarget = idx
	return
}

func (c *Config) TargetByAlias(alias string) (int, error) {
	for i, target := range c.data.Targets {
		if target.Alias == alias {
			return i, nil
		}
	}
	return -1, errors.New(fmt.Sprintf("Target given by alias='%s' not found.", alias))
}
