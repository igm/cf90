package main

import (
	"fmt"
)

type TargetList []*Target

func (t TargetList) Title() string     { return "Targets:" }
func (t TargetList) Selection() string { return "Select Target: " }
func (t TargetList) Len() int          { return len(t) }
func (t TargetList) Render(i int) string {
	return fmt.Sprintf("%-20s [%s - %s]", t[i].Alias, t[i].TargetUrl, t[i].Username)
}
