package main

import (
  "fmt"
  "os"
  "github.com/entropyx/yadet/cmd"
  "github.com/google/go-github/github"
)

func main() {
  if err := cmd.Yadet.Execute(); err != nil {
      fmt.Println(err)
      os.Exit(-1)
  }
  client := github.NewClient(nil)
  usr, _, err := client.Users.Get("arturodz")
  if err != nil {
    print(err)
  }
  fmt.Printf(usr.String());
}
