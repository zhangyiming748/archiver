package code

import (
	"log"
	"os"
	"path/filepath"

	"github.com/h2non/filetype"
	"github.com/zhangyiming748/archive"
)

func FindImageAndCovertImmediately(root string, threads int) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // 忽略错误，继续遍历
		}
		if !info.IsDir() {
			absPath, _ := filepath.Abs(path)
			if isImage(absPath) {
				if err := archive.Convert2AVIF(absPath, threads); err != nil {
					log.Printf("文件%s在处理的时候出现了错误:%v\n", absPath, err)
				}
			}
		}
		return nil
	})
}

func isImage(fp string) bool {
	// if strings.ToLower(filepath.Ext(fp)) == ".avif" {
	// 	// avif文件的magic number比较特殊，filetype库无法识别，所以直接通过后缀名判断
	// 	return true
	// }
	file, _ := os.Open(fp)
	defer file.Close()
	head := make([]byte, 261)
	file.Read(head)
	if filetype.IsImage(head) {
		return true
	} else {
		return false
	}
}
