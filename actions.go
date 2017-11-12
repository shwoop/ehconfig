package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

// ActionFunc defines the signature of the action functions.
type ActionFunction (func(Config) error)

// putText reads text from stdin and updates the objects config.
func putText(c Config) error {
	var input string
	var line []string

	f, err := claimLock(c)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := decodeConfigFile(c)
	if err != nil {
		return err
	}
	strm := bufio.NewScanner(os.Stdin)

	for strm.Scan() {
		input = strings.TrimRight(strm.Text(), "\n")
		line = strings.Split(input, " ")
		updateInfo(info, line[0], strings.Join(line[1:], " "))
	}
	if err != nil {
		printToStderr(err.Error())
		return err
	}

	writeToConfigFile(c, info)
	return nil
}

// putJson reads json envoded text from stdin and updats the objects config.
func putJson(c Config) error {
	input := make(JsonMapping)

	f, err := claimLock(c)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := decodeConfigFile(c)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(os.Stdin)

	err = dec.Decode(&input)
	if err != nil {
		printToStderr(err.Error())
		return err
	}

	for k := range input {
		updateInfo(info, k, input[k])
	}

	writeToConfigFile(c, info)
	return nil
}

// getSingleValue prints a value stored against the objects config.
func getSingleValue(c Config) error {
	info, err := decodeConfigFile(c)
	if err != nil {
		return err
	}
	if val := info[c.key]; val != "" {
		fmt.Println(val)
		return nil
	}
	return errors.New("value does not exist")
}

// getAll prints the entire config for a given object.
func getAll(c Config) error {
	info, err := decodeConfigFile(c)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(info)
	return nil
}
