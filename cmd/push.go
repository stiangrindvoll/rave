// Copyright © 2016 Stian Grindvoll <stian@grindvoll.org>

package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
	"github.com/stiangrindvoll/rave/discovery"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push a file or directory to your friend in the network",
	Long: `push will let you transfer a file or directory over to your friend
	in the network`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		ip, port := discovery.GetService(key)
		if ip == "" || port == "" {
			fmt.Fprintln(os.Stderr, "No Rave found")
			os.Exit(1)
		}

		if len(args) < 0 {
			fmt.Fprintln(os.Stderr, "Please select a file/directory you want to send")
			cmd.Usage()
			os.Exit(1)
		}

		conn, err := net.Dial("tcp", ip+":"+port)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unable to open tcp connection")
			os.Exit(1)
		}
		defer conn.Close()

		err = pushFiles(args, conn)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

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
