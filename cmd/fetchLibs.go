/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fetchLibsCmd represents the fetchLibs command
var fetchLibsCmd = &cobra.Command{
	Use:   "fetchLibs",
	Short: "fetch available repos from github/boostorg",
	Long:  `Fetch all available Boostorg repos from github`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("fetchLibs called")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(fetchLibsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchLibsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchLibsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
