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
	force     bool
	threads   int
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
			code.FindVideoAndCovertImmediately(rootDir, fhd, force)
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
			code.FindImageAndCovertImmediately(rootDir, threads)
			fmt.Printf("Image conversion completed for directory: %s\n", rootDir)
		},
	}
	// 创建旋转命令
	var rotateCmd = &cobra.Command{
		Use:   "rotate",
		Short: "Rotate video files",
		Long:  "Rotate all video files in specified directory",
		Run: func(cmd *cobra.Command, args []string) {
			dir, _ := cmd.Flags().GetString("dir")
			rotateDirection, _ := cmd.Flags().GetString("rotate")
			fmt.Printf("Starting video rotation task...\nDirectory: %s\nDirection: %s degrees\n", dir, rotateDirection)
			code.RotateVideos(dir, rotateDirection)
		},
	}

	// 创建 mp4 转换命令
	var mp4Cmd = &cobra.Command{
		Use:   "mp4",
		Short: "Convert video files to H265 MP4 format",
		Long:  "Find all video files in the specified directory and convert them to H265 MP4 format",
		Run: func(cmd *cobra.Command, args []string) {
			code.FindVideoAndCovertMp4Immediately(rootDir, fhd, force)
			fmt.Printf("MP4 conversion completed for directory: %s\n", rootDir)
		},
	}

	// 为 rotate 命令添加标志
	rotateCmd.Flags().StringP("dir", "d", "./", "Directory path for video rotation (required)")
	rotateCmd.Flags().StringP("rotate", "r", "90", "Rotation direction: 90, 270")
	// rotateCmd.MarkFlagRequired("dir")

	videoCmd.Flags().StringVarP(&rootDir, "dir", "d", "", "Directory path to search for video files")
	videoCmd.Flags().BoolVarP(&fhd, "fhd", "f", false, "Enable FHD mode for video conversion")
	videoCmd.Flags().BoolVar(&force, "force", false, "Force overwrite existing files")
	videoCmd.MarkFlagRequired("dir")

	mp4Cmd.Flags().StringVarP(&rootDir, "dir", "d", "", "Directory path to search for video files")
	mp4Cmd.Flags().BoolVarP(&fhd, "fhd", "f", false, "Enable FHD mode for MP4 conversion")
	mp4Cmd.Flags().BoolVar(&force, "force", false, "Force overwrite existing files")
	mp4Cmd.MarkFlagRequired("dir")

	imageCmd.Flags().StringVarP(&rootDir, "dir", "d", "", "Directory path to search for image files")
	imageCmd.Flags().BoolVarP(&fhd, "fhd", "f", false, "Enable FHD mode for image conversion")
	imageCmd.Flags().IntVarP(&threads, "threads", "t", 4, "Number of threads to use for conversion")
	imageCmd.MarkFlagRequired("dir")

	rootCmd.AddCommand(videoCmd)
	rootCmd.AddCommand(imageCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(rotateCmd)
	rootCmd.AddCommand(mp4Cmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
