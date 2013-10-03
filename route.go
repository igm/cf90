package main

import (
	"fmt"
	"github.com/igm/cf"
)

type RouteList []cf.Route

func (r RouteList) Len() int { return len(r) }
func (r RouteList) Render(i int) string {
	domain := fmt.Sprintf("%s.%s", r[i].Host, r[i].Domain.Name)
	return fmt.Sprintf("%-35s %-17s %s", domain, r[i].Space.Name, r[i].Apps)
}
func (r RouteList) Title() string     { return fmt.Sprintf("%-35s %-17s %s", "Routes", "Space", "Apps") }
func (r RouteList) Selection() string { return "Select route:" }
