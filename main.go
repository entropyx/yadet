package main

import (
  "github.com/google/go-github/github"
  "fmt"
)

func main() {
  client := github.NewClient(nil)
  usr, _, err := client.Users.Get("arturodz")
  if err != nil {
    print(err)
  }

  fmt.Printf(usr.String());
}
