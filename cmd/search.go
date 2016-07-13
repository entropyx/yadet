// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

  "github.com/google/go-github/github"
  // "github.com/entropyx/yadet/client"
  // "github.com/entropyx/yadet/cmd"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search developers",
	Long: `Search developers using several filters. For example:
    `,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
    client := github.NewClient(nil)
    users, _, err := client.Users.ListAll(nil)
    if err != nil {
      fmt.Println(err)
    }

    fmt.Print(users)
		fmt.Println("search called")
	},
}

func init() {
	Yadet.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
