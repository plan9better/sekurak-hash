package main

import (
	"os"
	"strings"
  "time"
)

// 300th line
const myhash = "257d84791c8008f78b93f0170a4ae15222fb28fd9ae60fc9d8ec0acddfa356a0"

func main() {
	bytes, _ := os.ReadFile("slowa.txt")
	str := string(bytes)
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		lines[i] = strings.Trim(line, "\r")
	}
  
  progress := 0
  pass := "";

  go func(){
    for pass == "" {
      println(progress, "/", len(lines))
      time.Sleep(time.Second)
    }
  }()
	pass = GetPassword(lines, myhash, &progress)
	println(pass)
}
