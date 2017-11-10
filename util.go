package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type JsonMapping map[string]string

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

func updateInfo(info JsonMapping, k, v string) {
	if v == "" {
		delete(info, k)
	} else {
		info[k] = v
	}
}

func printToStderr(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}

func checkError(err error) {
	if err != nil {
		printToStderr(err.Error())
		os.Exit(1)
	}
}

func decodeConfigFile() JsonMapping {
	output := make(JsonMapping)
	if _, err := os.Stat(config.configFile); !os.IsNotExist(err) {
		f, err := os.Open(config.configFile)
		checkError(err)
		defer f.Close()

		dec := json.NewDecoder(f)
		err = dec.Decode(&output)
		checkError(err)
	}
	return output
}

func writeToConfigFile(info JsonMapping) {
	err := os.MkdirAll(config.configDir, 0777)
	checkError(err)
	tmpFile, err := ioutil.TempFile(config.configDir, "")
	checkError(err)
	defer os.Remove(tmpFile.Name())

	enc := json.NewEncoder(tmpFile)
	enc.Encode(info)
	tmpFile.Close()
	err = os.Rename(tmpFile.Name(), config.configFile)
	checkError(err)
}
