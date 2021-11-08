package main

import "fmt"

func runHelp() {
	fmt.Print(`
qvault allows you to pass-off Qvault.io projects locally

  Find more information at https://qvault.io

Basic Commands:
  auth [key]     Create the config file ($home/.qvault.json) and populate it

`)

}
