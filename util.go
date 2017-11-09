package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func help(exitCode int) {
	fmt.Println("ehinfo json TYPE UUID")
	fmt.Println("       put  TYPE UUID")
	fmt.Println("       get  TYPE UUID KEY")
	fmt.Println("       info TYPE UUID")
	fmt.Println("")
	fmt.Println("Actions:")
	fmt.Println("  json: Place provided json onto object state")
	fmt.Println("  put:  Place key value data from stdin on object state")
	fmt.Println("  get:  Retreive value of key from object state")
	fmt.Println("  info: Retreive entire state from object")
	fmt.Println("")
	os.Exit(exitCode)
}

func checkType(objType string) (string, error) {
	objType = strings.ToLower(objType)
	switch objType {
	case "guest", "container", "drive", "folder":
	default:
		return "", errors.New("Unsupported type: " + objType)
	}
	return objType, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
