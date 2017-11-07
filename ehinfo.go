package main

import (
  "bytes"
  "fmt"
  "os"
  "bufio"
  "io"
  "strings"
  "encoding/json"
  "errors"
)

type ActionFunction func()

type Config struct {
  stateDir, lockDir, key, configFile, lockFile string
  action ActionFunction
}
var config Config

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

func checkType(objType string) (string, error) {
  objType = strings.ToLower(objType)
  switch objType {
  case "guest":
  case "container":
  case "drive":
  case "folder":
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
