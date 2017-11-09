package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type ActionFunction func()

func putText() {

	var line []string
	info := decodeConfigFile()
	strm := bufio.NewReader(os.Stdin)

	input, err := strm.ReadString('\n')
	for err != io.EOF {
		input = strings.TrimRight(input, "\n")
		line = strings.Split(input, " ")
		if len(line) == 1 {
			delete(info, line[0])
		} else if len(line) == 2 {
			info[line[0]] = line[1]
		}
		input, err = strm.ReadString('\n')
	}

	// Original code -- delete once no longer a reference
	// // Read stdin into map
	// var line []string
	// strm := bufio.NewReader(os.Stdin)
	// output := make(map[string]string)
	// input, err := strm.ReadString('\n')
	// for err != io.EOF {
	// 	input = strings.TrimRight(input, "\n")
	// 	line = strings.Split(input, " ")
	// 	if len(line) == 1 {
	// 		delete(output, line[0])
	// 	} else if len(line) == 2 {
	// 		output[line[0]] = line[1]
	// 	}
	// 	input, err = strm.ReadString('\n')
	// }

	// // Encode map into json and print to stdout
	// enc := json.NewEncoder(os.Stdout)
	// enc.Encode(output)
	// fmt.Println(config.stateDir)
}

func putJson() {
}

func getSingleValue() {
	info := decodeConfigFile()
	if val := info[config.key]; val != "" {
		fmt.Println(val)
		os.Exit(0)
	}
	os.Exit(1)
}

func decodeConfigFile() map[string]string {
	if _, err := os.Stat(config.configFile); os.IsNotExist(err) {
		os.Exit(1)
	}
	f, err := os.Open(config.configFile)
	checkError(err)
	defer f.Close()

	dec := json.NewDecoder(f)
	output := make(map[string]string)
	err = dec.Decode(&output)
	checkError(err)

	return output
}

func writeToConfigFile() {
	tmpFile, err := ioutil.TempFile(config.configDir, "")
	defer os.RemoveFile(tempFile.Name())
}

func getAll() {
	info := decodeConfigFile()
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(info)
}
