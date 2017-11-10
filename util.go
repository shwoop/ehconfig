package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// JsonMapping is the dictionary like mapping used to store an objects config.
type JsonMapping map[string]string

// help outputs helpful information about calling the program.
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

// checkType validates the provided object type is known and supported.
func checkType(objType string) (string, error) {
	objType = strings.ToLower(objType)
	switch objType {
	case "guest", "container", "drive", "folder":
	default:
		return "", errors.New("Unsupported type: " + objType)
	}
	return objType, nil
}

// updateInfo manages updating an objects config in memory.
// If a value is presented it is updated/edited on the config, otherwise the
// value is removed.
func updateInfo(info JsonMapping, k, v string) {
	if v == "" {
		delete(info, k)
	} else {
		info[k] = v
	}
}

// printTostderr is a helper funciton to print a string to stderr.
func printToStderr(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}

// checkError checks if an error has been triggered.
// In the case of an error, it's message is printed to stderr and the program
// is closed.
func checkError(err error) {
	if err != nil {
		printToStderr(err.Error())
		os.Exit(1)
	}
}

// decodeConfigFile returns the existing config for an object.
// It attempts to retreive a config from the filesystem, in the event one does
// not exist an empty JsonMapping is returned.
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

// writeToConfigFile writes a provided JsonMapping to the filesystem.
// The function is atomic, in that it writes to a temporary file and swaps them
// rather than writing sequentially to the config file.
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
