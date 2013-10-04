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

func (s SpaceList) FindOrPick(space, org string) (i int, err error) {
	i, ok := find(s, func(i int) bool { return s[i].Name == space && s[i].Organization.Name == org })
	if !ok {
		i, err = choose(s)
		if err != nil {
			return -1, err
		}
	}
	return i, nil
}

func (target *Target) SpaceFind(name, org string) (space cf.Space, err error) {
	// Get Application ID
	spaces, err := target.SpacesGet()
	if err != nil {
		return
	}
	i, err := SpaceList(spaces).FindOrPick(name, org)
	if err != nil {
		return
	}
	space = spaces[i]
	return
}
