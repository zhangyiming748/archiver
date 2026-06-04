package main

import (
	"fmt"
	"os"

	"archiver/code"

	"github.com/spf13/cobra"
)

var (
	version = "1.0.0"
	rootDir string
	fhd     bool
)

func main() {
	var rootCmd = &cobra.Command{
		Use:     "archiver",
		Short:   "Archiver is a CLI tool for media file management",
		Long:    "Archiver is a command line tool for managing and converting media files",
		Version: version,
	}

	var videoCmd = &cobra.Command{
		Use:   "video",
		Short: "Convert video files to H265 format",
		Long:  "Find all video files in the specified directory and convert them to H265 format",
		Run: func(cmd *cobra.Command, args []string) {
			code.FindAndCovertImmediately(rootDir, fhd)
			fmt.Printf("Video conversion completed for directory: %s\n", rootDir)
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("archiver version %s\n", version)
		},
	}

	videoCmd.Flags().StringVarP(&rootDir, "dir", "d", "", "Directory path to search for video files")
	videoCmd.Flags().BoolVarP(&fhd, "fhd", "f", false, "Enable FHD mode for video conversion")
	videoCmd.MarkFlagRequired("dir")

	rootCmd.AddCommand(videoCmd)
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
