package pathutils

import (
	"mime"
	"path/filepath"
	"strings"
)

func ContentType(path string) string {
	ext := strings.ToLower(filepath.Ext(path))

	ctype := ""

	if ext != "" {
		ctype = MimeMap[ext[1:]]
	}

	if ctype == "" {
		ctypeParts := strings.Split(mime.TypeByExtension(ext), ";")
		ctype = ctypeParts[0]
	}

	if ctype == "" {
		ctype = "application/octet-stream"
	}

	return ctype
}
