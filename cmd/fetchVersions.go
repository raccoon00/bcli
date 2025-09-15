/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

type ReleaseInfo struct {
	Tag     string `json:"tag_name"`
	Created string `json:"created_at"`
}

func fetchReleases() ([]ReleaseInfo, error) {
	url := "https://api.github.com/repos/boostorg/boost/releases?page=1"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error making GET request to github server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Returned status code by api.github.com/releases is not OK")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error when reading body of a response of api.github.com/releases: %w", err)
	}

	var releases []ReleaseInfo
	err = json.Unmarshal(body, &releases)
	if err != nil {
		return nil, fmt.Errorf("Error when unmarshaling api.github.com/releases json, %w", err)
	}

	return releases, nil
}

// fetchVersionsCmd represents the fetchVersion command
var fetchVersionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "Fetch available boost versions with dates",
	Long:  `Fetch available boost releases from api.github.com`,
	RunE: func(cmd *cobra.Command, args []string) error {
		releases, err := fetchReleases()
		if err != nil {
			return fmt.Errorf("Error fetching releases: %w", err)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		for _, release := range releases {
			fmt.Fprintf(w, "%v\t%v\n", release.Tag, release.Created)
		}
		w.Flush()

		return nil
	},
}

func init() {
	fetchCmd.AddCommand(fetchVersionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchVersionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchVersionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
