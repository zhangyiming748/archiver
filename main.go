package main

import (
	"fmt"
	"os"

	"archiver/code"

	"github.com/spf13/cobra"
)

var (
	version   = "dev"
	buildTime = "unknown"
	gitCommit = "unknown"
	rootDir   string
	fhd       bool
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
			code.FindVideoAndCovertImmediately(rootDir, fhd)
			fmt.Printf("Video conversion completed for directory: %s\n", rootDir)
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("archiver version %s\n", version)
			fmt.Printf("Build time: %s\n", buildTime)
			fmt.Printf("Git commit: %s\n", gitCommit)
		},
	}

	var imageCmd = &cobra.Command{
		Use:   "image",
		Short: "Convert image files to AVIF format",
		Long:  "Find all image files in the specified directory and convert them to AVIF format",
		Run: func(cmd *cobra.Command, args []string) {
			code.FindImageAndCovertImmediately(rootDir)
			fmt.Printf("Image conversion completed for directory: %s\n", rootDir)
		},
	}

	videoCmd.Flags().StringVarP(&rootDir, "dir", "d", "", "Directory path to search for video files")
	videoCmd.Flags().BoolVarP(&fhd, "fhd", "f", false, "Enable FHD mode for video conversion")
	videoCmd.MarkFlagRequired("dir")

	imageCmd.Flags().StringVarP(&rootDir, "dir", "d", "", "Directory path to search for image files")
	imageCmd.Flags().BoolVarP(&fhd, "fhd", "f", false, "Enable FHD mode for image conversion")
	imageCmd.MarkFlagRequired("dir")

	rootCmd.AddCommand(videoCmd)
	rootCmd.AddCommand(imageCmd)
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
