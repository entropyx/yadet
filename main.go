package main

import (
  "fmt"
  "os"
  "github.com/entropyx/yadet/cmd"
  "github.com/google/go-github/github"
)

func main() {
  if err := cmd.RootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }
}
