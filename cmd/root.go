/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

type ChromeBookmarks struct {
	Roots struct {
		BookmarkBar Child `json:"bookmark_bar"`
	} `json:"roots"`
}

type Child struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Url      string  `json:"url"`
	Children []Child `json:"children"`
	Path     string
}

func ScrubPath(s string) string {
	// todo dedeupe "_"
	re, err := regexp.Compile(`[^\w/]`)
	if err != nil {
		log.Fatal(err)
	}
	return strings.ToLower(re.ReplaceAllString(s, "_"))
}

func Walk(c *Child, path string) error {
	c.Path = ScrubPath(fmt.Sprintf("%s/%s", c.Path, c.Name))
	if len(c.Children) == 0 {
		return nil
	} else {
		for i, _ := range c.Children {
			Walk(&c.Children[i], c.Path)
		}
	}

	return nil
}

// todo print of bookmarks (flat, nested?)
// todo create bookmark files

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

		Walk(&bm.Roots.BookmarkBar, "")

		fmt.Printf("%+v\n", bm)
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
