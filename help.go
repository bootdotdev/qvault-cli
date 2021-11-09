package main

import "fmt"

func runHelp() {
	fmt.Print(`
qvault allows you to pass-off Qvault.io projects locally

  Find more information at https://qvault.io

Basic Commands:
  auth [key]     Create the config file ($HOME/.qvault.json) and set the auth credentials
  project [uuid] Create the local config file (.qvault.json) and set the project id
  check          Checks your code, passes it off if it is correct
  help           Prints this guide
`)
}
