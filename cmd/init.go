/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init boost in this repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil

		os.Mkdir("boost", 0755)

		gitrepo := "https://github.com/boostorg/boost.git"
		oscmd := exec.Command("git", "clone", "--depth", "1", gitrepo)
		err := oscmd.Run()
		if err != nil {
			return fmt.Errorf("Error when cloning boost git repo: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
