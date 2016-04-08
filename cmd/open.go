// Copyright © 2016 Stian Grindvoll <stian@grindvoll.org>

package cmd

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/stiangrindvoll/rave/discovery"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a directory to recieve files in a directory from a network friend.",
	Long:  `Open a directory to recieve files in a directory from a network friend.`,

	Run: func(cmd *cobra.Command, args []string) {

		disc, err := discovery.New("mDNS", "Rave Client", "Rave File Sharing")
		if err != nil {
			panic(err)
		}
		s, err := disc.Register()
		defer s.Shutdown()

		if err != nil {
			panic(err)
		}
		log.Fatal(http.ListenAndServe(":1623", http.FileServer(http.Dir("./"))))
	},
}

func init() {
	RootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
