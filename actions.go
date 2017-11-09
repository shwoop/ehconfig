package main

import (
  "strings"
  "fmt"
  "bufio"
  "encoding/json"
  "io"
  "os"
)

type ActionFunction func()

func putText() {
  var line []string
  strm := bufio.NewReader(os.Stdin)
  output := make(map[string]string)
  input, err := strm.ReadString('\n')
  for err != io.EOF {
    input = strings.TrimRight(input, "\n")
    line = strings.Split(input, " ")
    output[line[0]] = line[1]
    input, err = strm.ReadString('\n')
  }
  enc := json.NewEncoder(os.Stdout)
  enc.Encode(output)
  fmt.Println(config.stateDir)
}

func putJson() {
}

func getSingleValue() {
}

func getAll() {
}
