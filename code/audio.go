package code

import (
	"github.com/zhangyiming748/archive"
	"github.com/zhangyiming748/finder"
)

func Novel(dir string) {
	fs := finder.FindAllAudios(dir)
	for _, f := range fs {
		archive.ConvertAudio(f, archive.AudioBookType)
	}
}
