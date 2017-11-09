package main

import (
	"bytes"
	"fmt"
	"os"
)

type Config struct {
	stateDir, lockDir, key, configFile, lockFile string
	action                                       ActionFunction
}

var config Config

func init() {
	args := os.Args[1:]
	if len(args) == 0 {
		help(0)
	}
	var objType string
	var objUuid string
	var err error
	config = Config{}
	argLength := len(args)
	if argLength >= 3 {
		config.action = putJson
		objType, err = checkType(args[1])
		checkError(err)
		objUuid = args[2]
	}
	if argLength == 3 {
		switch args[0] {
		case "json":
			config.action = putJson
		case "put":
			config.action = putText
		case "info":
			config.action = getAll
		default:
			help(1)
		}
	} else if argLength == 4 {
		if args[0] != "get" {
			help(1)
		}
		config.action = getSingleValue
		config.key = args[3]
	}

	config.stateDir = os.Getenv("STATEDIR")
	config.lockDir = os.Getenv("LOCKDIR")
	// testing
	config.stateDir = "/tmp/eh/state"
	config.lockDir = "/tmp/eh/lock"

	var filename bytes.Buffer
	filename.WriteString(config.stateDir)
	filename.WriteRune('/')
	filename.WriteString(objType)
	filename.WriteRune('/')
	filename.WriteString(objUuid)
	filename.WriteString("/config.json")
	config.configFile = filename.String()

	filename.Reset()

	filename.WriteString(config.lockDir)
	filename.WriteRune('/')
	filename.WriteString(objType)
	filename.WriteRune('/')
	filename.WriteString(objUuid)
	config.lockFile = filename.String()
}

func main() {
	config.action()
	fmt.Println("lock file", config.lockFile)
	fmt.Println("config file", config.configFile)
}
