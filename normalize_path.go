package pathutils

import (
	"strings"
)

func NormalizePath(path string) (newPath string, ok bool) {
	path = replaceMultiSlashes(path)

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

func replaceMultiSlashes(path string) string {
	// regex is slow
	// BenchmarkReplaceMultiSlashesRegexp-8     2000000               735 ns/op
	// BenchmarkReplaceMultiSlashesRunes-8     10000000               236 ns/op
	// BenchmarkReplaceMultiSlashesBytes-8     20000000               82.1 ns/op
	pathBytes := []byte(path)
	bs := make([]byte, 0, len(pathBytes))
	isLastSlash := false
	for _, b := range pathBytes {
		isSlash := b == '/'
		if isSlash && isLastSlash {
			continue
		}
		isLastSlash = isSlash
		bs = append(bs, b)
	}
	return string(bs)
}
