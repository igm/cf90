package main

import (
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
	apps, err := target.AppsGet()
	if err != nil {
		log.Fatal(err)
	}

	name, path, appId := params["name"], params["path"], ""
	for _, app := range apps {
		if app.Name == name {
			appId = app.Guid
		}
	}

	if appId == "" {
		i, err := choose(AppList(apps))
		if err != nil {
			log.Fatal(err)
		}
		appId = apps[i].Guid
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
