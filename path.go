package pathutils

import (
	"path/filepath"
)

func SplitNameExt(fileName string) (name, ext string) {
	ext = filepath.Ext(fileName)

	if len(ext) > 0 {
		name = fileName[:len(fileName)-len(ext)]
		ext = ext[1:] // remove dot
	} else {
		name = fileName
	}

	return
}

func JoinNameExt(name, ext string) (fileName string) {
	fileName = name

	if len(ext) > 0 {
		fileName += "." + ext
	}

	return
}
