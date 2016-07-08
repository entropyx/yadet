package main

import (
  "fmt"
  "os"
  "github.com/entropyx/yadet/cmd"
)

func main() {
    if err := cmd.Yadet.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}
