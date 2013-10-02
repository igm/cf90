package main

import (
	"encoding/json"
	"errors"
	"github.com/igm/cf"
	"os"
)

type Config struct {
	filename string
	trace    bool
	data     struct {
		ActiveTarget int          `json:"selected"`
		ActiveOrg    string       `json:"org"`
		ActiveSpace  string       `json:"space"`
		Targets      []*cf.Target `json:"targets"`
	}
}

func NewConfig(filename string) (ret *Config, err error) {
	ret = new(Config)
	ret.filename = filename
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&ret.data)
	return
}

func (c *Config) Save() (err error) {
	file, err := os.Create(c.filename)
	if err != nil {
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(&c.data)
	return
}

func (c *Config) AddTarget(target *cf.Target) {
	for _, existing := range c.data.Targets {
		if existing.TargetUrl == target.TargetUrl {
			return
		}
	}
	c.data.Targets = append(c.data.Targets, target)
}

func (c *Config) SelectedTarget() (*cf.Target, error) {
	if c.data.ActiveTarget >= len(c.data.Targets) {
		return nil, errors.New("Not target selected.")
	}
	ret := c.data.Targets[c.data.ActiveTarget]
	return ret, nil
}

func (c *Config) Select(host string) error {
	for i, existing := range c.data.Targets {
		if existing.TargetUrl == host {
			c.data.ActiveTarget = i
			return nil
		}
	}
	return errors.New("Target does not  exist")
}

func (c *Config) SelectedOrg() (string, error) {
	if c.data.ActiveOrg == "" {
		return c.data.ActiveOrg, errors.New("No organization selected.")
	}
	return c.data.ActiveOrg, nil
}
