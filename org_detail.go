package main

import (
	"log"
	"os"
	"text/template"
)

// space.use also sets org
func init() {
	register(&Command{
		group:  "Organization",
		name:   "org.detail",
		help:   "Show organization detail",
		params: []Param{Param{name: "org", desc: "Organization name"}},
		handle: org_detail,
	})
}

func org_detail() {
	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}

	orgs, err := target.OrganizationsGet()
	if err != nil {
		log.Fatal(err)
	}

	i, err := OrgList(orgs).findOrg(params["org"])
	if err != nil {
		i, err = choose(OrgList(orgs))
		if err != nil {
			log.Fatal(err)
		}
	}
	orgTmpl.Execute(os.Stdout, orgs[i])
	return
}

var orgTmpl = template.Must(template.New("orgDetail").Parse(org))

const org = `
Organization: {{.Name}} ({{.Status}})

Spaces:
{{range .Spaces}}    {{.Name}} 
{{end}}
Domains:
{{range .Domains}}    {{.Name}} ( spaces:{{range .Spaces}} {{.Name}} {{end}})
{{end}}
`
