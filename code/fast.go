package code

import (
	"os"
	"path/filepath"
	"github.com/zhangyiming748/archive"
)

func FindVideoAndCovertSmallerMp4Immediately(root string, fhd, force bool) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // 忽略错误，继续遍历
		}
		if !info.IsDir() {
			absPath, _ := filepath.Abs(path)
			if isVideo(absPath) {
				archive.Convert2SmallerH265MP4(absPath, fhd, force)
			}
		}
		return nil
	})
}


