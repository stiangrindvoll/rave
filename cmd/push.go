// Copyright © 2016 Stian Grindvoll <stian@grindvoll.org>

package cmd

import (
	"archive/tar"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"

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
			os.Exit(1)
		}
		conn, err := net.Dial("tcp", ip+":"+port)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unable to open tcp connection")
			os.Exit(1)
		}
		defer conn.Close()
		tarWriter := tar.NewWriter(conn)
		defer tarWriter.Close()

		for _, f := range args {
			info, err := os.Stat(f)
			if err != nil {
				os.Exit(1)
			}
			var baseDir string
			if info.IsDir() {
				baseDir = filepath.Base(f)
			}

			filepath.Walk(f,
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}

					header, err := tar.FileInfoHeader(info, info.Name())
					if err != nil {
						return err
					}

					if baseDir != "" {
						header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, f))
					}

					if err = tarWriter.WriteHeader(header); err != nil {
						return err
					}

					if info.IsDir() {
						return nil
					}

					file, err := os.Open(path)
					if err != nil {
						return err
					}
					defer file.Close()
					_, err = io.Copy(tarWriter, file)
					return err

				})
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
