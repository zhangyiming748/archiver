// Package main 是 archiver 项目的入口点
// archiver 是一个用于媒体文件管理的命令行工具，支持视频转码、图片转换等功能
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/zhangyiming748/archive"
	"github.com/zhangyiming748/archiver/code"
)

// 全局变量定义
var (
	version   = "dev"     // 版本号，构建时通过 ldflags 注入
	buildTime = "unknown" // 构建时间，构建时通过 ldflags 注入
	gitCommit = "unknown" // Git 提交哈希，构建时通过 ldflags 注入
	rootDir   string      // 根目录路径，用于指定要处理的文件目录
	fhd       bool        // FHD 模式标志，启用全高清视频处理
	force     bool        // 强制覆盖标志，是否覆盖已存在的文件
	threads   int         // 线程数，用于控制并行处理的线程数量
)

// ========== 初始化函数 ==========

// init 在项目启动时初始化必要的依赖
// 例如:数据库连接、配置文件加载等
func init() {
	// 初始化 SQLite 数据库(用于处理进度记录)
	archive.InitSqlte()
}

// ========== 命令创建区域 ==========

func main() {
	// 创建根命令
	var rootCmd = &cobra.Command{
		Use:     "archiver",
		Short:   "Archiver is a CLI tool for media file management",
		Long:    "Archiver is a command line tool for managing and converting media files",
		Version: version,
	}

	// ---------- 子命令定义 ----------

	// version 命令：打印版本信息
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("archiver version %s\n", version)
			fmt.Printf("Build time: %s\n", buildTime)
			fmt.Printf("Git commit: %s\n", gitCommit)
		},
	}

	// video 命令：将视频文件转换为 H265 格式
	var videoCmd = &cobra.Command{
		Use:   "video",
		Short: "Convert video files to H265 format",
		Long:  "Find all video files in the specified directory and convert them to H265 format",
		Run: func(cmd *cobra.Command, args []string) {
			code.FindVideoAndCovertImmediately(rootDir, fhd, force)
			fmt.Printf("Video conversion completed for directory: %s\n", rootDir)
		},
	}

	// mp4 命令：将视频文件转换为 H265 MP4 格式
	var mp4Cmd = &cobra.Command{
		Use:   "mp4",
		Short: "Convert video files to H265 MP4 format",
		Long:  "Find all video files in the specified directory and convert them to H265 MP4 format",
		Run: func(cmd *cobra.Command, args []string) {
			code.FindVideoAndCovertMp4Immediately(rootDir, fhd, force)
			fmt.Printf("MP4 conversion completed for directory: %s\n", rootDir)
		},
	}

	// smart 命令：智能转换视频文件为更小的 H265 MP4 格式
	var smartCmd = &cobra.Command{
		Use:   "smart",
		Short: "Smart convert video files to smaller H265 MP4 format",
		Long:  "Find all video files in the specified directory and smart convert them to smaller H265 MP4 format",
		Run: func(cmd *cobra.Command, args []string) {
			code.FindVideoAndCovertSmallerMp4Immediately(rootDir, fhd, force)
			fmt.Printf("Smart MP4 conversion completed for directory: %s\n", rootDir)
		},
	}

	// rotate 命令：旋转视频文件
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

	// image 命令:将图片文件转换为 AVIF 格式
	var imageCmd = &cobra.Command{
		Use:   "image",
		Short: "Convert image files to AVIF format",
		Long:  "Find all image files in the specified directory and convert them to AVIF format",
		Run: func(cmd *cobra.Command, args []string) {
			code.FindImageAndCovertImmediately(rootDir, threads)
			fmt.Printf("Image conversion completed for directory: %s\n", rootDir)
		},
	}

	// novel 命令：转换音频文件为有声小说格式
	var novelCmd = &cobra.Command{
		Use:   "novel",
		Short: "Convert audio files to audiobook novel format",
		Long:  "Find all audio files in the specified directory and convert them to audiobook novel format",
		Run: func(cmd *cobra.Command, args []string) {
			code.Novel(rootDir)
			fmt.Printf("Novel audiobook conversion completed for directory: %s\n", rootDir)
		},
	}

	// opus 命令：将音频文件转换为 Opus 格式
	var opusCmd = &cobra.Command{
		Use:   "opus",
		Short: "Convert audio files to Opus format",
		Long:  "Find all audio files in the specified directory and convert them to Opus format",
		Run: func(cmd *cobra.Command, args []string) {
			code.Opus(rootDir)
			fmt.Printf("Opus conversion completed for directory: %s\n", rootDir)
		},
	}

	// ---------- 参数配置 ----------

	// version 命令无需额外参数

	// video 命令参数
	videoCmd.Flags().StringVarP(&rootDir, "dir", "d", "", "Directory path to search for video files")
	videoCmd.Flags().BoolVarP(&fhd, "fhd", "f", false, "Enable FHD mode for video conversion")
	videoCmd.Flags().BoolVar(&force, "force", false, "Force overwrite existing files")
	videoCmd.MarkFlagRequired("dir")

	// mp4 命令参数
	mp4Cmd.Flags().StringVarP(&rootDir, "dir", "d", "", "Directory path to search for video files")
	mp4Cmd.Flags().BoolVarP(&fhd, "fhd", "f", false, "Enable FHD mode for MP4 conversion")
	mp4Cmd.Flags().BoolVar(&force, "force", false, "Force overwrite existing files")
	mp4Cmd.MarkFlagRequired("dir")

	// smart 命令参数
	smartCmd.Flags().StringVarP(&rootDir, "dir", "d", "", "Directory path to search for video files")
	smartCmd.Flags().BoolVarP(&fhd, "fhd", "f", false, "Enable FHD mode for smart MP4 conversion")
	smartCmd.Flags().BoolVar(&force, "force", false, "Force overwrite existing files")
	smartCmd.MarkFlagRequired("dir")

	// rotate 命令参数
	rotateCmd.Flags().StringP("dir", "d", "./", "Directory path for video rotation")
	rotateCmd.Flags().StringP("rotate", "r", "90", "Rotation direction: 90, 270")

	// image 命令参数
	imageCmd.Flags().StringVarP(&rootDir, "dir", "d", "", "Directory path to search for image files")
	imageCmd.Flags().BoolVarP(&fhd, "fhd", "f", false, "Enable FHD mode for image conversion")
	imageCmd.Flags().IntVarP(&threads, "threads", "t", 4, "Number of threads to use for conversion")
	imageCmd.MarkFlagRequired("dir")

	// novel 命令参数
	novelCmd.Flags().StringVarP(&rootDir, "dir", "d", ".", "Directory path to search for audio files")

	// opus 命令参数
	opusCmd.Flags().StringVarP(&rootDir, "dir", "d", ".", "Directory path to search for audio files")

	// ---------- 注册命令 ----------
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(videoCmd)
	rootCmd.AddCommand(mp4Cmd)
	rootCmd.AddCommand(smartCmd)
	rootCmd.AddCommand(rotateCmd)
	rootCmd.AddCommand(imageCmd)
	rootCmd.AddCommand(novelCmd)
	rootCmd.AddCommand(opusCmd)

	// ---------- 执行命令 ----------
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
