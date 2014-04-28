package pathutils_test

import (
	. "git.koofr.lan/go-pathutils.git"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Escape", func() {
	Describe("EscapePath", func() {
		It("should escape path parts and keep /", func() {
			Expect(EscapePath("path/to/file?name.txt")).To(Equal("path/to/file%3Fname.txt"))
		})
	})
})
