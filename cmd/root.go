/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type ChromeBookmarks struct {
	Roots struct {
		BookmarkBar Child `json:"bookmark_bar"`
	} `json:"roots"`
}

type Child struct {
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Url         string      `json:"url"`
	ChildrenRaw interface{} `json:"children"`
	Children    []Child
	Path        string
}

var rootCmd = &cobra.Command{
	Use:   "bm-cli",
	Short: "Parse and flatten bookmarks while maintaining file structure",
	Run: func(cmd *cobra.Command, args []string) {
		fn := "/Users/andrew.cooper/Library/Application Support/Google/Chrome/Default/Bookmarks"
		file, err := os.ReadFile(fn)
		if err != nil {
			fmt.Println(err)
		}

		var bm ChromeBookmarks
		err = json.Unmarshal(file, &bm)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(bm)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
