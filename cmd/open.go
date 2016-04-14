// Copyright © 2016 Stian Grindvoll <stian@grindvoll.org>

package cmd

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/stiangrindvoll/rave/discovery"
)

var workpath string

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a directory to recieve files in a directory from a network friend.",
	Long:  `Open a directory to recieve files in a directory from a network friend.`,

	Run: func(cmd *cobra.Command, args []string) {

		disc, err := discovery.New("mDNS", key, "Rave File Sharing")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		err = disc.Register()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer disc.Close()

		l, err := net.Listen("tcp", ":1623")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer l.Close()

		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer conn.Close()

		workdir, err := filepath.Abs(workpath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		err = writeFiles(conn, workdir)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	},
}

func init() {
	RootCmd.AddCommand(openCmd)

	openCmd.PersistentFlags().StringVarP(&workpath, "path", "p", filepath.Dir(os.Args[0]), "Path to where files should be saved")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
