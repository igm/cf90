/*
	Cloud Foundry/Pivotal CF command line client

*/
package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/igm/cf"
	"log"
	"net/http"
	"os/user"
	"sort"
	"strings"
)

type (
	Param struct {
		name   string
		desc   string
		defval string
	}
	Command struct {
		group  string
		name   string
		help   string
		params []Param
		handle func()
	}
	Commands []*Command
)

var c *Config
var commands Commands
var params map[string]string = make(map[string]string)

func register(cmd *Command)      { commands = append(commands, cmd) }
func (c Commands) Len() int      { return len(c) }
func (c Commands) Swap(a, b int) { c[a], c[b] = c[b], c[a] }
func (c Commands) Less(a, b int) bool {
	if c[a].group == c[b].group {
		return c[a].name < c[b].name
	}
	return c[a].group < c[b].group
}

func init() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configFile := fmt.Sprintf("%s/.cf90", user.HomeDir)
	c, _ = NewConfig(configFile)

	flag.BoolVar(&cf.Trace, "t", false, "Enable HTTP tracing")
	insecure := flag.Bool("k", false, "This option explicitly allows cf90 to perform \"insecure\" SSL connections and transfers.")

	flag.Parse()
	log.SetFlags(0)

	for _, arg := range flag.Args() {
		parsed := strings.Split(arg, "=")
		if len(parsed) == 2 {
			params[parsed[0]] = parsed[1]
		}
	}

	// use more toleran HTTP Client
	if *insecure {
		cf.HttpClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				Proxy:           http.ProxyFromEnvironment,
			},
			CheckRedirect: http.DefaultClient.CheckRedirect,
		}
	}
}

func main() {
	sort.Sort(commands)
	defer c.Save()

	for _, command := range commands {
		if command.name == flag.Arg(0) {
			// fill in default values
			for _, param := range command.params {
				if params[param.name] == "" && param.defval != "" {
					params[param.name] = param.defval
				}
			}
			command.handle()
			return
		}
	}
	help()
}
