package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func init() {
	register(&Command{
		group: "Application",
		name:  "app.fetch",
		help:  "Fetch application instance files into local zip file",
		params: []Param{
			Param{name: "name", desc: "Application name"},
			Param{name: "org", desc: "Organization name"},
			Param{name: "space", desc: "Space name"},
			Param{name: "instance", desc: "Application instance", defval: "0"},
			Param{name: "dir", desc: "Remote dir"},
		},
		handle: app_fetch,
	})
}

func app_fetch() {

	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	// Get Application ID
	app, err := target.AppFind(params["name"], params["space"], params["org"])
	if err != nil {
		log.Fatal(err)
	}

	instance := params["instance"]
	timestamp := time.Now().Format("20060102150405")
	filename := fmt.Sprintf("%s.%s.%s.zip", app.Name, instance, timestamp)
	zipOut, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("creating zip: %s\n", filename)
	zipWriter := zip.NewWriter(zipOut)
	defer zipWriter.Close()

	var fetchDir func(string)

	fetchDir = func(dir string) {
		files, err := target.AppLs(app.Guid, instance, dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			if file.Dir {
				fmt.Printf("recuring into %s%s\n", dir, file.Name)
				fetchDir(dir + file.Name)
			} else {
				remoteFile := fmt.Sprintf("%s%s", dir, file.Name)
				fmt.Printf("fetching %s\n", remoteFile)
				reader, err := target.AppGet(app.Guid, instance, remoteFile)
				if err != nil {
					log.Fatal(err)
				}
				w, err := zipWriter.Create(remoteFile)
				if err != nil {
					log.Fatal(err)
				}
				io.Copy(w, reader)
			}
		}
	}
	fetchDir(params["dir"] + "/")
}
