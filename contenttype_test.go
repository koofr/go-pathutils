package pathutils_test

import (
	. "github.com/koofr/go-pathutils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ContentType", func() {
	It("should return content type for text file", func() {
		Expect(ContentType("path/to/file.txt")).To(Equal("text/plain"))
		Expect(ContentType("path/to/file.m4v")).To(Equal("video/x-m4v"))
	})

	It("should return content type for file without ext", func() {
		Expect(ContentType("path/to/file")).To(Equal("application/octet-stream"))
	})

	It("should return content type for file with empty name", func() {
		Expect(ContentType("")).To(Equal("application/octet-stream"))
	})
})
