package pathutils_test

import (
	. "github.com/koofr/go-pathutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ContentType", func() {
	It("should return content type for text file", func() {
		Expect(ContentType("path/to/file.txt")).To(Equal("text/plain"))
	})

	It("should return content type for file without ext", func() {
		Expect(ContentType("path/to/file")).To(Equal("application/octet-stream"))
	})

	It("should return content type for file with empty name", func() {
		Expect(ContentType("")).To(Equal("application/octet-stream"))
	})
})
