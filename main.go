package main

import "os"

// main is the entrypoint fo the application.
func main() {
	c, err := buildConfig()
	if err == nil || c.action(c) != nil {
		os.Exit(1)
	}
}
