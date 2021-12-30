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
		urls, err := find(args)
		if err != nil {
			panic(err)
		}
		for _, url := range urls {
			fmt.Println(url)
		}
	},
}

var OpenCmd = &cobra.Command{
	Use:   "open",
	Short: "Find the url associated with bookmark id",
	Run: func(cmd *cobra.Command, args []string) {
		urls, err := find(args)
		if err != nil {
			panic(err)
		}
		for _, url := range urls {
			fmt.Println(url)
		}
	},
}

func find(args []string) ([]string, error) {
	var ids []string
	var rawId []string

	if len(args) > 0 {
		for _, raw := range args {
			rawId = strings.Split(raw, ":")
			ids = append(ids, rawId[0])
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			rawId = strings.Split(scanner.Text(), ":")
			ids = append(ids, rawId[0])
		}
	}

	urls, err := cbm.Find(path, ids)
	if err != nil {
		return ids, nil
	}

	return urls, nil
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
