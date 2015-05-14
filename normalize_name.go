package pathutils

import (
	"bytes"
	"golang.org/x/text/unicode/norm"
)

func NormalizeName(name string) string {
	n := string(norm.NFC.Bytes([]byte(name)))

	var buffer bytes.Buffer
	for _, r := range n {
		if r >= '\u0020' && !(r >= '\u007f' && r <= '\u009f') {
			buffer.WriteRune(r)
		}
	}
	n = buffer.String()

	return n
}
