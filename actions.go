package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// ActionFunc defines the signature of the action functions.
type ActionFunction func(Config)

// putText reads text from stdin and updates the objects config.
func putText(c Config) {
	var input string
	var line []string

	f := claimLock(c)
	defer f.Close()

	info := decodeConfigFile(c)
	strm := bufio.NewScanner(os.Stdin)

	for strm.Scan() {
		input = strings.TrimRight(strm.Text(), "\n")
		line = strings.Split(input, " ")
		updateInfo(info, line[0], strings.Join(line[1:], " "))
	}
	checkError(strm.Err())

	writeToConfigFile(c, info)
}

// putJson reads json envoded text from stdin and updats the objects config.
func putJson(c Config) {
	input := make(JsonMapping)

	f := claimLock(c)
	defer f.Close()

	info := decodeConfigFile(c)
	dec := json.NewDecoder(os.Stdin)

	err := dec.Decode(&input)
	checkError(err)

	for k := range input {
		updateInfo(info, k, input[k])
	}

	writeToConfigFile(c, info)
}

// getSingleValue prints a value stored against the objects config.
func getSingleValue(c Config) {
	info := decodeConfigFile(c)
	if val := info[c.key]; val != "" {
		fmt.Println(val)
		os.Exit(0)
	}
	os.Exit(1)
}

// getAll prints the entire config for a given object.
func getAll(c Config) {
	info := decodeConfigFile(c)
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(info)
}
