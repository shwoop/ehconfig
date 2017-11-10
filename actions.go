package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"os"
	"strings"
)

type ActionFunction func()

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

func getSingleValue() {
	info := decodeConfigFile()
	if val := info[config.key]; val != "" {
		fmt.Println(val)
		os.Exit(0)
	}
	os.Exit(1)
}

func getAll() {
	info := decodeConfigFile()
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(info)
}
