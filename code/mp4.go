package code

import (
	"os"
	"path/filepath"

	"github.com/zhangyiming748/archive"
)

func FindVideoAndCovertMp4Immediately(root string, fhd, force bool, limit int) {
	count := 0
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // 忽略错误，继续遍历
		}
		if !info.IsDir() {
			absPath, _ := filepath.Abs(path)
			if isVideo(absPath) {
				archive.Convert2H265MP4(absPath, fhd, force)
				count++
				// limit 大于 0 时限制处理数量，等于 0 时处理所有视频
				if limit > 0 && count >= limit {
					return filepath.SkipAll
				}
			}
		}
		return nil
	})
}

func FindVideoAndCovertHEVCImmediately(root string, fhd, force bool, limit int) {
	count := 0
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // 忽略错误，继续遍历
		}
		if !info.IsDir() {
			absPath, _ := filepath.Abs(path)
			if isVideo(absPath) {
				archive.ReEncode2H265KeepContainer(absPath, fhd, force)
				count++
				// limit 大于 0 时限制处理数量，等于 0 时处理所有视频
				if limit > 0 && count >= limit {
					return filepath.SkipAll
				}
			}
		}
		return nil
	})
}
