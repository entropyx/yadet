package main

import (
  "github.com/google/go-github/github"
  "fmt"
)

func main() {
  if err := cmd.RootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }
}
