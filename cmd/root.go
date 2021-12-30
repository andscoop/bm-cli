/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/andscoop/bm-cli/cbm"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(FindCmd)
}

var rootCmd = &cobra.Command{
	Use:   "cbm",
	Short: "Reads reads input from stdin and treats it like find",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("try --help")
	},
}

var ListCmd = &cobra.Command{
	Use:   "ls",
	Short: "Output bookmarks in flattened list, suitable for fzf",
	Run: func(cmd *cobra.Command, args []string) {
		err := cbm.FlatList(path)
		if err != nil {
			panic(err)
		}
	},
}

var path string = "/Users/andrew.cooper/Library/Application Support/Google/Chrome/Default/Bookmarks"

var FindCmd = &cobra.Command{
	Use:   "find",
	Short: "Find the url associated with bookmark id",
	Run: func(cmd *cobra.Command, args []string) {
		var parts []string

		if len(args) > 0 {
			parts = strings.Split(args[0], ":")
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				parts = strings.Split(scanner.Text(), ":")
			}
		}

		url, err := cbm.Find(path, parts[0])
		if err != nil {
			panic(err)
		}

		fmt.Println(url)
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
