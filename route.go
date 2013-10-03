package main

import (
	"fmt"
	"github.com/igm/cf"
)

type RouteList []cf.Route

func (r RouteList) Len() int            { return len(r) }
func (r RouteList) Render(i int) string { return fmt.Sprintf("%s.%s", r[i].Host, r[i].Domain.Name) }
func (r RouteList) Title() string       { return "Routes" }
func (r RouteList) Selection() string   { return "Select route:" }
