package main

import (
	"fmt"
	"github.com/igm/cf"
)

type AppList []cf.App

func (a AppList) Len() int { return len(a) }
func (a AppList) Title() string {
	return fmt.Sprintf("%-25s%-12s%-15s%-30s%-13s", "Name", "Status", "Usage", "SpaceOrg", "Routes")
}
func (a AppList) Render(i int) string {
	instances := fmt.Sprintf("%d x %dM", a[i].Instances, a[i].Memory)
	return fmt.Sprintf("%-25s%-12s%-15s%-30s%-13s", a[i].Name, a[i].State, instances, a[i].Space, a[i].Routes)
}
func (a AppList) Selection() string { return "Select application: " }

func (a AppList) FindOrPick(name, space, org string) (i int, err error) {
	i, ok := find(a, func(i int) bool {
		return a[i].Name == name && a[i].Space.Name == space && a[i].Space.Organization.Name == org
	})
	if !ok {
		i, err = choose(a)
		if err != nil {
			return -1, err
		}
	}
	return i, nil
}

func (target *Target) AppFind(name, space, org string) (app cf.App, err error) {
	// Get Application ID
	apps, err := target.AppsGet()
	if err != nil {
		return
	}
	i, err := AppList(apps).FindOrPick(name, space, org)
	if err != nil {
		return
	}
	app = apps[i]
	return
}
