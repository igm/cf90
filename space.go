package main

import (
	"fmt"
	"github.com/igm/cf"
)

type SpaceList []cf.Space

func (s SpaceList) Len() int { return len(s) }
func (s SpaceList) Render(i int) string {
	return fmt.Sprintf("%s", s[i])
}
func (s SpaceList) Title() string     { return fmt.Sprintf("%s", "SpaceOrg") }
func (s SpaceList) Selection() string { return "Select space:" }
