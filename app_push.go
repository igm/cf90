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
		help:  "Push application (directory or file archive)",
		params: []Param{
			Param{name: "name", desc: "Application name"},
			Param{name: "path", desc: "Application path, or path to a archive file to upload (zip, war,...)", defval: "."},
		},
		handle: app_push,
	})
}

func app_push() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}

	// Get Application ID
	app, err := target.AppFind(params["name"], params["space"], params["org"])
	if err != nil {
		log.Fatal(err)
	}

	path := params["path"]
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if stat.IsDir() {
		var archetypes []*cf.Archetype
		filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				file, err := os.Open(p)
				if err != nil {
					return err
				}
				fmt.Println("Adding file:", info.Name())
				norm, _ := filepath.Rel(path, p)
				archetypes = append(archetypes, &cf.Archetype{
					Name:   norm,
					Reader: file,
				})
			}
			return nil
		})
		err = target.AppPush(app.Guid, archetypes)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Pushing file:", stat.Name())
		err = target.AppPushArchive(app.Guid, file)
	}
}
