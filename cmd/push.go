// Copyright © 2016 Stian Grindvoll <stian@grindvoll.org>

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
			fmt.Println("No Rave found")
			os.Exit(1)
		}
		res, err := http.Get(fmt.Sprintf("http://%v:%v/index.html", ip, port))
		if err != nil {
			log.Fatal(err)
		}
		robots, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", robots)

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
