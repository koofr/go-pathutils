package pathutils

import (
	"strings"
)

func IsPathValid(path string) bool {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	if strings.Contains(path, "\\") { // backslash
		return false
	}
	if strings.Contains(path, "//") { // two slashes
		return false
	}
	if strings.Contains(path, "/./") { // dot
		return false
	}
	if strings.Contains(path, "/../") { // two dots
		return false
	}

	return true
}

func IsNameValid(name string) bool {
	if name == "" {
		return false
	}
	if name == "." {
		return false
	}
	if name == ".." {
		return false
	}
	if strings.Contains(name, "\\") { // backslash
		return false
	}
	if strings.Contains(name, "/") { // slash
		return false
	}

	return true
}
