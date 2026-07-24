package code

import (
	"github.com/zhangyiming748/finder"
	"github.com/zhangyiming748/archive"
	
)
func Novel(dir string) {
	fs := finder.FindAllAudios(dir)
	for _, f := range fs {
		archive.ConvertAudio(f, archive.AudioBookType)
	}
}
