package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	register(&Command{
		name:   "help",
		help:   "Shows this help message, use [COMMAND] for command parameters",
		handle: help,
	})
}

func help() {
	for _, command := range commands {
		if command.name == flag.Arg(1) {
			if len(command.params) > 0 {
				fmt.Printf("%s [options] %s [param1=value1 param2=value2 ...]\n", os.Args[0], command.name)
			} else {
				fmt.Printf("%s [options] %s\n", os.Args[0], command.name)
			}
			fmt.Println("\nOptions:")
			flag.PrintDefaults()
			fmt.Println("\nDescription:\n ", command.help)
			if len(command.params) > 0 {
				fmt.Print("\nParameters: ")
				for _, param := range command.params {
					fmt.Printf("\n  %-13s  %s", param.name, param.desc)
					if param.defval != "" {
						fmt.Printf(", default value: %s", param.defval)
					}
				}
			}
			fmt.Printf("\n\n")
			return
		}
	}

	fmt.Printf("%s [options] command [param1=value1 param2=value2 ...]", os.Args[0])
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
	fmt.Println("\nCommands:")
	group := ""
	for _, cmd := range commands {
		if cmd.group != group {
			fmt.Printf("%s\n", cmd.group)
			group = cmd.group
		}
		fmt.Printf("  %-15s - %s\n", cmd.name, cmd.help)
	}
	fmt.Println("")
}
