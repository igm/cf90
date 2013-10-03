package main

import (
	"fmt"
	"github.com/igm/cf"
)

type DomainList []cf.Domain

func (d DomainList) Len() int            { return len(d) }
func (d DomainList) Title() string       { return fmt.Sprintf("%-20s  %s", "Domain", "Spaces") }
func (d DomainList) Render(i int) string { return fmt.Sprintf("%-20s  %s", d[i].Name, d[i].Spaces) }
func (d DomainList) Selection() string   { return "Select domain:" }
