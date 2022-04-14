package main

import (
	"os"

	"github.com/Cyber-SiKu/cli/cmd"
)

func main() {
  error := cmd.Execute()
  if error != nil {
	os.Exit(-1)
  }
}