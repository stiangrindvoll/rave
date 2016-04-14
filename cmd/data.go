package cmd

import (
	"archive/tar"
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func pushFiles(files []string, conn net.Conn) error {
	tarWriter := tar.NewWriter(conn)
	defer tarWriter.Close()

	for _, f := range files {
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
	return nil
}

func writeFiles(conn net.Conn, workdir string) error {
	buf := make([]byte, 1024)
	tarReader := tar.NewReader(conn)
	for {
		hdr, err := tarReader.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}
		if err != nil {
			return err
		}
		if hdr.FileInfo().IsDir() {
			err = os.MkdirAll(path.Join(workdir, hdr.Name), hdr.FileInfo().Mode())
			if err != nil {
				return err
			}
		}
		if hdr.FileInfo().Mode().IsRegular() {
			file, err := os.OpenFile(path.Join(workdir, hdr.Name), os.O_CREATE|os.O_WRONLY, hdr.FileInfo().Mode())
			if err != nil {
				return err
			}

			_, err = io.CopyBuffer(file, tarReader, buf)
			if err != nil {
				return err
			}

			file.Close()
		}

		fmt.Println("Recived:", hdr.Name)
	}

	return nil
}
