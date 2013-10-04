package main

import (
	"fmt"
	"github.com/igm/cf"
)

type Target struct {
	Alias string `json:"alias"`
	// Org       string `json:"org"`
	// OrgGuid   string `json:"org_guid"`
	// Space     string `json:"space"`
	// SpaceGuid string `json:"space_guid"`
	*cf.Target
}

type TargetList []*Target

func (t TargetList) Title() string     { return "Targets:" }
func (t TargetList) Selection() string { return "Select Target: " }
func (t TargetList) Len() int          { return len(t) }
func (t TargetList) Render(i int) string {
	return fmt.Sprintf("%-20s [%s - %s]", t[i].Alias, t[i].TargetUrl, t[i].Username)
}
