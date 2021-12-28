/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/andscoop/bm-cli/cbm"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ListCmd)
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

func lsCmd(cmd *cobra.Command, args []string) {
	path := "/Users/andrew.cooper/Library/Application Support/Google/Chrome/Default/Bookmarks"
	err := cbm.FlatList(path)
	if err != nil {
		panic(err)
	}

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
