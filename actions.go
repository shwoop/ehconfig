package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// ActionFunc defines the signature of the action functions.
type ActionFunction func()

// putText reads text from stdin and updates the objects config.
func putText() {
	var input string
	var line []string
	info := decodeConfigFile()
	strm := bufio.NewScanner(os.Stdin)

	for strm.Scan() {
		input = strings.TrimRight(strm.Text(), "\n")
		line = strings.Split(input, " ")
		updateInfo(info, line[0], strings.Join(line[1:], " "))
	}
	checkError(strm.Err())

	writeToConfigFile(info)
}

// putJson reads json envoded text from stdin and updats the objects config.
func putJson() {
	input := make(JsonMapping)
	info := decodeConfigFile()
	dec := json.NewDecoder(os.Stdin)

	err := dec.Decode(&input)
	checkError(err)

	for k := range input {
		updateInfo(info, k, input[k])
	}

	writeToConfigFile(info)
}

// getSingleValue prints a value stored against the objects config.
func getSingleValue() {
	info := decodeConfigFile()
	if val := info[config.key]; val != "" {
		fmt.Println(val)
		os.Exit(0)
	}
	os.Exit(1)
}

// getAll prints the entire config for a given object.
func getAll() {
	info := decodeConfigFile()
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(info)
}
