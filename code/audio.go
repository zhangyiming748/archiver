package code

import (
	"fmt"
	"path/filepath"

	"github.com/zhangyiming748/archive"
	"github.com/zhangyiming748/finder"
)

func Novel(dir string) {
	fs := finder.FindAllAudios(dir)
	for _, f := range fs {
		if filepath.Ext(f)==".opus"{
			fmt.Println("Skipping opus file:", f)
			continue
		}
		archive.ConvertAudio(f, archive.AudioBookType)
	}
}

func Opus(dir string) {
	fs := finder.FindAllAudios(dir)
	for _, f := range fs {
		if filepath.Ext(f)==".opus"{
			fmt.Println("Skipping opus file:", f)
			continue
		}
		archive.Convert2Opus(f)
	}
}
