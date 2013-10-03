package main

import (
	"fmt"
	"github.com/igm/cf"
)

type OrgList []cf.Organization

func (o OrgList) Len() int            { return len(o) }
func (o OrgList) Render(i int) string { return fmt.Sprintf("%s", o[i].Name) }
func (o OrgList) Title() string       { return "Organizations" }
func (o OrgList) Selection() string   { return "Select organization:" }
