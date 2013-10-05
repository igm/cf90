package main

import (
	"fmt"
	"github.com/igm/cf"
)

type AppList []cf.App

func (a AppList) Len() int { return len(a) }
func (a AppList) Title() string {
	return fmt.Sprintf("%-20s%-12s%-15s%-30s%-13s", "Name", "Status", "Usage", "SpaceOrg", "Routes")
}
func (a AppList) Render(i int) string {
	instances := fmt.Sprintf("%d x %dM", a[i].Instances, a[i].Memory)
	return fmt.Sprintf("%-20s%-12s%-15s%-30s%-13s", a[i].Name, a[i].State, instances, a[i].Space, a[i].Routes)
}
func (a AppList) Selection() string { return "Select application: " }

func (a AppList) FindOrPick(name, space, org string) (app cf.App, err error) {

	filtered := a.
		Filter(func(app cf.App) bool { return app.Name == name || name == "" }).
		Filter(func(app cf.App) bool { return app.Space.Name == space || space == "" }).
		Filter(func(app cf.App) bool { return app.Space.Organization.Name == org || org == "" })

	if len(filtered) == 0 {
		filtered = a
	}
	if len(filtered) != 1 {
		i, err := choose(filtered)
		if err != nil {
			return filtered[0], err
		}
		app = filtered[i]
		return app, err
	}
	return filtered[0], nil
}

func (a AppList) Filter(fn func(cf.App) bool) AppList {
	var res AppList
	for _, val := range a {
		if fn(val) {
			res = append(res, val)
		}
	}
	return res
}

func (target *Target) AppFind(name, space, org string) (app cf.App, err error) {
	apps, err := target.AppsGet()
	if err != nil {
		return
	}
	app, err = AppList(apps).FindOrPick(name, space, org)
	return
}
