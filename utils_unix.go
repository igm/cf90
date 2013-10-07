// +build linux darwin freebsd

package main

func bold(str string) string {
	return "\033[1m" + str + "\033[0m"
}
