package pathutils

import (
	"strings"
)

func NormalizePath(path string) (newPath string, ok bool) {
	if !IsPathValid(path) {
		return "", false
	}

	if path == "" || path == "/" {
		return "/", true
	}

	if path[0] == '/' {
		path = path[1:]
	}

	if path[len(path)-1] == '/' {
		path = path[0 : len(path)-1]
	}

	parts := strings.Split(path, "/")

	for i, part := range parts {
		newPart := NormalizeName(part)

		if !IsNameValid(newPart) {
			return "", false
		}

		parts[i] = newPart
	}

	newPath = "/" + strings.Join(parts, "/")

	return newPath, true
}
