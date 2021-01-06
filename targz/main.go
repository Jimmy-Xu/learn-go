package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	targzAbsPath string
)

func main() {
	rootCmd := &cobra.Command{
		Use: "app",
	}
	//tar subcommand
	var (
		tarSource string
		tarTarget string
	)
	tarCmd := &cobra.Command{
		Use:   "tar",
		Short: "tar file",
		Long:  "tar file",
		Run: func(cmd *cobra.Command, args []string) {
			if f, err := os.Open(tarSource); err != nil {
				log.Fatalf("failed to open tar source %v, error:%v", tarSource, err)
			} else {
				if err = Compress(f, tarTarget); err != nil {
					log.Fatalf("failed to tar file %v to $%v, error:%v", tarSource, tarTarget, err)
				}
			}
			log.Printf("tar file %v to %v OK", tarSource, tarTarget)
		},
	}
	tarCmd.Flags().StringVarP(&tarSource, "source", "s", ".", "source file or dir")
	tarCmd.Flags().StringVarP(&tarTarget, "target", "t", "", "target file")

	//untar subcommand
	var (
		untarSource string
		untarTarget string
	)
	untarCmd := &cobra.Command{
		Use:   "untar",
		Short: "untar file",
		Long:  "untar file",
		Run: func(cmd *cobra.Command, args []string) {
			if err := DeCompress(untarSource, untarTarget); err != nil {
				log.Fatalf("failed to untar file %v from $%v, error:%v", untarTarget, untarSource, err)
			}
			log.Printf("untar file %v from %v OK", untarTarget, untarSource)
		},
	}
	untarCmd.Flags().StringVarP(&untarSource, "source", "s", "", "source file")
	untarCmd.Flags().StringVarP(&untarTarget, "target", "t", ".", "target dir")

	//add subcommand
	rootCmd.AddCommand(tarCmd)
	rootCmd.AddCommand(untarCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

//压缩 使用gzip压缩成tar.gz
func Compress(file *os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()

	targzAbsPath, _ = filepath.Abs(dest)
	log.Printf("abs path of dest(%v): %v", dest, targzAbsPath)

	gw := gzip.NewWriter(d)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	err := compress(file, "", tw)
	if err != nil {
		return err
	}
	return nil
}

func compress(file *os.File, prefix string, tw *tar.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			absPath, _ := filepath.Abs(fmt.Sprintf("%v/%v", file.Name(), fi.Name()))
			if absPath == targzAbsPath {
				log.Printf("current file is dest: %v, skip", absPath)
				continue
			}
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, tw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := tar.FileInfoHeader(info, "")
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//解压 tar.gz
func DeCompress(tarFile, dest string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := dest + hdr.Name
		file, err := createFile(filename)
		if err != nil {
			return err
		}
		io.Copy(file, tr)
	}
	return nil
}

func createFile(name string) (*os.File, error) {
	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}
