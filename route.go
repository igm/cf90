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

func (r RouteList) FindOrPick(host, domain string) (i int, err error) {
	i, ok := find(r, func(i int) bool {
		return r[i].Host == host && r[i].Domain.Name == domain
	})
	if !ok {
		i, err = choose(r)
		if err != nil {
			return -1, err
		}
	}
	return i, nil
}

func (target *Target) RouteFind(host, domain string) (route cf.Route, err error) {
	routes, err := target.RoutesGet()
	if err != nil {
		return
	}
	i, err := RouteList(routes).FindOrPick(host, domain)
	if err != nil {
		return
	}
	route = routes[i]
	return

}
