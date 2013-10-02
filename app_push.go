package main

import (
	"fmt"
	"github.com/igm/cf"
	"log"
	"os"
	"path/filepath"
)

func init() {
	register(&Command{
		group: "Application",
		name:  "app.push",
		help:  "Push application",
		params: []Param{
			Param{name: "name", desc: "Application name"},
			Param{name: "path", desc: "Application path", defval: "."},
		},
		handle: app_push,
	})
}

func app_push() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	summary, err := target.Summary(c.data.ActiveSpace)
	if err != nil {
		log.Fatal(err)
	}

	name, path, appId := params["name"], params["path"], ""
	for _, app := range summary.Apps {
		if app.Name == name {
			appId = app.Guid
		}
	}

	if appId == "" {
		name, appId = chooseApplication(summary)
	}

	var archetypes []*cf.Archetype
	filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			file, err := os.Open(p)
			if err != nil {
				return err
			}
			norm, _ := filepath.Rel(path, p)
			archetypes = append(archetypes, &cf.Archetype{
				Name:   norm,
				Reader: file,
			})
		}
		return nil
	})
	err = target.AppPush(appId, archetypes)
	if err != nil {
		log.Fatal(err)
	}
}

func chooseApplication(summary *cf.Summary) (string, string) {
	for i, app := range summary.Apps {
		fmt.Printf(" (%d) %s\n", i+1, app.Name)
	}
	var appIndex int
	fmt.Print("Select application: ")
	_, err := fmt.Scanf("%d\n", &appIndex)
	if err != nil || appIndex > len(summary.Apps) || appIndex < 1 {
		return "", ""
	}
	return summary.Apps[appIndex-1].Name, summary.Apps[appIndex-1].Guid
}
