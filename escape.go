package pathutils

import (
	"net/url"
	"strings"
)

func EscapePath(path string) string {
	parts := strings.Split(path, "/")

	for i := range parts {
		parts[i] = url.QueryEscape(parts[i])
	}

	return strings.Join(parts, "/")
}
