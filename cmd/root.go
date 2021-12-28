/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/andscoop/bm-cli/cbm"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(FindCmd)
}

var rootCmd = &cobra.Command{
	Use:   "cbm",
	Short: "Aliased to 'ls'",
	Run:   lsCmd,
}

var ListCmd = &cobra.Command{
	Use:   "ls",
	Short: "Output bookmarks in flattened list, suitable for fzf",
	Run:   lsCmd,
}

var path string = "/Users/andrew.cooper/Library/Application Support/Google/Chrome/Default/Bookmarks"

func lsCmd(cmd *cobra.Command, args []string) {
	err := cbm.FlatList(path)
	if err != nil {
		panic(err)
	}
}

var FindCmd = &cobra.Command{
	Use:   "find",
	Short: "Find the url associated with bookmark id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cbm.Find(path, args[0])
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
