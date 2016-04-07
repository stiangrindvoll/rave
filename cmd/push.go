// Copyright © 2016 Stian Grindvoll <stian@grindvoll.org>

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stiangrindvoll/rave/discovery"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push a file or directory from your network friend.",
	Long: `push will let you push down an available file or directory
	from your network friend`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		discovery.List()

	},
}

func init() {
	RootCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
