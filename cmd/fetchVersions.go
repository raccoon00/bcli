package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var fetchVersionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "Fetch available boost versions with dates",
	Long:  `Fetch available boost releases from api.github.com`,
	RunE: func(cmd *cobra.Command, args []string) error {
		releases, err := FetchReleases()
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
