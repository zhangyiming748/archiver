package code

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
	"github.com/zhangyiming748/archive"
)

func FindVideoAndCovertImmediately(root string, fhd, force bool) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // 忽略错误，继续遍历
		}
		if !info.IsDir() {
			absPath, _ := filepath.Abs(path)
			if isVideo(absPath) {
				archive.Convert2H265(absPath, fhd, force)
			}
		}
		return nil
	})
}

func isVideo(fp string) bool {
	file, _ := os.Open(fp)
	defer file.Close()
	head := make([]byte, 261)
	file.Read(head)
	ext := strings.ToLower(filepath.Ext(fp))
	if filetype.IsVideo(head) {
		return true
	} else if strings.ToLower(ext) == ".rmvb" {
		return true
	} else if strings.ToLower(ext) == ".rm" {
		return true
	} else if strings.ToLower(ext) == ".vob" {
		return true
	} else if strings.ToLower(ext) == ".flv" {
		return true
	} else if strings.ToLower(ext) == ".ts" {
		return true
	} else if strings.ToLower(ext) == ".m2ts" {
		return true
	} else {
		return false
	}
}

const (
	ToRight = "ClockWise90"
	ToLeft  = "ClockWise270"
)

func RotateVideos(root, direction string) {
	var archiveDirection string

	switch direction {
	case "90":
		archiveDirection = archive.ToRight
	case "270":
		archiveDirection = archive.ToLeft
	default:
		log.Printf("警告：无效的旋转方向 '%s'，使用默认方向 90 度\n", direction)
		archiveDirection = archive.ToRight
		direction = "90"
	}

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // 忽略错误，继续遍历
		}
		if !info.IsDir() {
			absPath, _ := filepath.Abs(path)
			if isVideo(absPath) {
				archive.RotateVideo(absPath,archiveDirection)
			}
		}
		return nil
	})
}

func RotateVideo(src string, direction string) {
	var (
		cmd  *exec.Cmd
		args []string
	)
	tmp_name := strings.Replace(src, filepath.Ext(src), "_rotate.mp4", 1)
	args = append(args, "-i", src)
	switch direction {
	case ToRight:
		args = append(args, "-vf", "transpose=1")
	case ToLeft:
		args = append(args, "-vf", "transpose=2")
	default:
		log.Printf("请输入正确的旋转方向:%s\n", direction)
		return
	}
	if hasNvidia() {
		// NVIDIA NVENC: 使用高质量预设，CRF模式保持画质
		args = append(args, "-c:v", "h264_nvenc")
		args = append(args, "-preset", "p5") // slow预设，质量与速度平衡
		args = append(args, "-rc", "vbr")    // 可变比特率
		args = append(args, "-cq", "18")     // 恒定质量等级，18为高质量
		args = append(args, "-b:v", "0")     // 不限制最大比特率
	} else if hasIntel() {
		// Intel QSV: 使用ICQ模式获得最佳质量
		args = append(args, "-c:v", "h264_qsv")
		args = append(args, "-global_quality", "18")   // ICQ质量等级，18为高质量
		args = append(args, "-look_ahead", "1")        // 启用前瞻分析
		args = append(args, "-look_ahead_depth", "40") // 前瞻深度
	} else if hasAMD() {
		// AMD AMF: 使用质量优先预设
		args = append(args, "-c:v", "h264_amf")
		args = append(args, "-quality", "quality") // 质量优先模式
		args = append(args, "-qp_i", "18")         // I帧量化参数
		args = append(args, "-qp_p", "20")         // P帧量化参数
		args = append(args, "-qp_b", "22")         // B帧量化参数
	} else {
		// CPU软件编码 libx264: 使用slow预设和CRF 18
		args = append(args, "-c:v", "libx264")
		args = append(args, "-preset", "slow")     // slow预设，质量与压缩率平衡
		args = append(args, "-crf", "18")          // CRF 18，视觉无损级别
		args = append(args, "-pix_fmt", "yuv420p") // 标准像素格式
	}
	args = append(args, "-tag:v", "avc1")
	args = append(args, "-c:a", "aac")
	args = append(args, "-map_chapters", "-1")
	args = append(args, tmp_name)
	cmd = exec.Command("ffmpeg", args...)
	log.Printf("开始执行命令:%s\n", cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("旋转失败：%v\n输出内容%s\n", err, string(out))
	} else {
		log.Printf("旋转成功：%s\n", string(out))
		/*
			1. 删除旧文件
			2. 临时文件改为旧文件的文件名
		*/
		if err := os.Remove(src); err != nil {
			log.Printf("删除源文件失败：%v\n", err)
		} else {
			if err := os.Rename(tmp_name, strings.Replace(src, filepath.Ext(src), ".mp4", 1)); err != nil {
				log.Printf("重命名文件失败：%v\n", err)
			}
		}
	}
}


func hasNvidia() bool {
	// 检查FFmpeg是否支持NVIDIA NVENC H.264编码器
	cmd := exec.Command("ffmpeg", "-encoders")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	// 查找h264_nvenc编码器
	return strings.Contains(string(output), "h264_nvenc")
}

func hasIntel() bool {
	// 检查FFmpeg是否支持Intel QSV H.264编码器
	cmd := exec.Command("ffmpeg", "-encoders")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	// 查找h264_qsv编码器
	return strings.Contains(string(output), "h264_qsv")
}

func hasAMD() bool {
	// 检查FFmpeg是否支持AMD VCE H.264编码器
	cmd := exec.Command("ffmpeg", "-encoders")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	// 查找h264_amf编码器
	return strings.Contains(string(output), "h264_amf")
}
