package client

import (
  "github.com/google/go-github/github"

)

// C Github Client
func C() *github.Client {
  client := github.NewClient(nil)
  return client
}
