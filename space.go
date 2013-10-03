package main

import (
	"fmt"
	"github.com/igm/cf"
)

type SpaceList []cf.Space

func (s SpaceList) Len() int { return len(s) }
func (s SpaceList) Render(i int) string {
	return fmt.Sprintf("%-20s %s", s[i].Name, s[i].Organization.Name)
}
func (s SpaceList) Title() string     { return fmt.Sprintf("%-20s %s", "Space", "Org") }
func (s SpaceList) Selection() string { return "Select space:" }
