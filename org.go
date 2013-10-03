package main

import (
	"fmt"
	"github.com/igm/cf"
)

type OrgList []cf.Organization

func (o OrgList) Len() int            { return len(o) }
func (o OrgList) Render(i int) string { return fmt.Sprintf("%-20s %s", o[i].Name, o[i].Domains) }
func (o OrgList) Title() string       { return fmt.Sprintf("%-20s %s", "Org", "Domains") }
func (o OrgList) Selection() string   { return "Select organization:" }
