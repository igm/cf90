package main

import (
	"fmt"
	"github.com/igm/cf"
)

type ServiceList []cf.Service

func (s ServiceList) Len() int { return len(s) }
func (s ServiceList) Title() string {
	return fmt.Sprintf("%-15s %-9s %-15s %s", "Service", "Version", "Provider", "Description")
}

func (s ServiceList) Render(i int) string {
	service := s[i]
	return fmt.Sprintf("%-15s %-9s %-15s %s", service.Label, service.Version, service.Provider, service.Description)
}
