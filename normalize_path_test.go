package pathutils_test

import (
	. "github.com/koofr/go-pathutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NormalizePath", func() {
	var ExpectPath = func(path string, expectedNewPath string) {
		newPath, ok := NormalizePath(path)
		Expect(newPath).To(Equal(expectedNewPath))
		Expect(ok).To(Equal(true))
	}

	var ExpectPathFail = func(path string) {
		_, ok := NormalizePath(path)
		Expect(ok).To(Equal(false))
	}

	It("should normalize path", func() {
		ExpectPath("", "/")
		ExpectPath("/", "/")
		ExpectPath("/foo", "/foo")
		ExpectPath("/foo/", "/foo")
		ExpectPath("/s\u030c", "/\u0161")
		ExpectPath("/\u0161", "/\u0161")
	})

	It("should fail to normalize path", func() {
		ExpectPathFail("//")
		ExpectPathFail("/\u0002")
		ExpectPathFail("/\u0002/foo.txt")
	})
})
