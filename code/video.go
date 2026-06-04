package code

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
	"github.com/zhangyiming748/archive"
)

func FindVideoAndCovertImmediately(root string, fhd bool) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // 忽略错误，继续遍历
		}
		if !info.IsDir() {
			absPath, _ := filepath.Abs(path)
			if isVideo(absPath) {
				archive.Convert2H265(absPath, fhd)
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
