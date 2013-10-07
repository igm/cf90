package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	register(&Command{
		group: "Application",
		name:  "app.start",
		help:  "Start application",
		params: []Param{
			Param{name: "name", desc: "Application name"},
			Param{name: "space", desc: "Space name"},
			Param{name: "org", desc: "Organization name"},
		},
		handle: app_start,
	})
}

func app_start() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	// Get Application ID
	app, err := target.AppFind(params["name"], params["space"], params["org"])
	if err != nil {
		log.Fatal(err)
	}

	resp, err := target.AppStart(app.Guid)
	if err != nil {
		log.Fatal(err)
	}
	stagingUrl := resp.Header.Get("X-App-Staging-Log")
	if stagingUrl != "" {
		readStaging(stagingUrl)
	}
}

func readStaging(url string) {
	written := 0
	for {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode == http.StatusNotFound {
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		if len(body) > written {
			io.Copy(os.Stdout, bytes.NewReader(body[written:]))
			written = len(body)
		}
		resp.Body.Close()
		<-time.After(1 * time.Second)
	}
}
