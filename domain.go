package main

import (
	"fmt"
	"github.com/igm/cf"
)

type DomainList []cf.Domain

func (d DomainList) Len() int            { return len(d) }
func (d DomainList) Title() string       { return fmt.Sprintf("Domain") }
func (d DomainList) Render(i int) string { return fmt.Sprintf("%s", d[i].Name) }
func (d DomainList) Selection() string   { return "Select domain:" }
