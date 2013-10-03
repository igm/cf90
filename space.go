package main

import (
	"fmt"
	"github.com/igm/cf"
)

type SpaceList []cf.Space

func (s SpaceList) Len() int            { return len(s) }
func (s SpaceList) Render(i int) string { return fmt.Sprintf("%s", s[i].Name) }
func (s SpaceList) Title() string       { return "Spaces" }
func (s SpaceList) Selection() string   { return "Select space:" }
