package main

import (
	"fmt"
	"github.com/igm/cf"
)

type AppList []cf.App

func (a AppList) Len() int { return len(a) }
func (a AppList) Title() string {
	return fmt.Sprintf("%-25s%-12s%-11s%-13s", "Name", "Status", "Usage", "URL")
}
func (a AppList) Render(i int) string {
	instances := fmt.Sprintf("%d x %dM", a[i].Instances, a[i].Memory)
	return fmt.Sprintf("%-25s%-12s%-11s%-13s", a[i].Name, a[i].State, instances, a[i].Urls)
}
func (a AppList) Selection() string { return "Select application: " }
