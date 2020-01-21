package pathutils_test

import (
	. "github.com/koofr/go-pathutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NormalizePath", func() {
	var ExpectPath = func(path string, expectedNewPath string) {
		newPath, ok := NormalizePath(path)
		ExpectWithOffset(1, newPath).To(Equal(expectedNewPath))
		ExpectWithOffset(1, ok).To(Equal(true))
	}

	var ExpectPathFail = func(path string) {
		_, ok := NormalizePath(path)
		ExpectWithOffset(1, ok).To(Equal(false))
	}

	It("should normalize path", func() {
		ExpectPath("", "/")
		ExpectPath("/", "/")
		ExpectPath("/foo", "/foo")
		ExpectPath("/foo/", "/foo")
		ExpectPath("~/a/b/c.txt", "/~/a/b/c.txt")
		ExpectPath("a.txt", "/a.txt")
		ExpectPath("a/b/c", "/a/b/c")
		ExpectPath("a/b/c/", "/a/b/c")
		ExpectPath("C:", "/C:")
		ExpectPath("~", "/~")
		ExpectPath("~/", "/~")
		ExpectPath("~user", "/~user")
		ExpectPath("~user/", "/~user")
		ExpectPath("/s\u030c", "/\u0161")
		ExpectPath("/\u0161", "/\u0161")
		ExpectPath("//server/foo/bar", "/server/foo/bar")
		ExpectPath("//server/bar//", "/server/bar")
		ExpectPath("/////", "/")
		ExpectPath("/////foo", "/foo")
		ExpectPath("foo/////", "/foo")
	})

	It("should fail to normalize path", func() {
		ExpectPathFail("/\u0002/foo.txt")
		ExpectPathFail("/\u0002/foo.txt")
		ExpectPathFail("/foo/../bar")
		ExpectPathFail("/foo/../bar/")
		ExpectPathFail("/foo/../bar/../baz")
		ExpectPathFail("/../")
		ExpectPathFail("../foo")
		ExpectPathFail("foo/bar/..")
		ExpectPathFail("foo/../../bar")
		ExpectPathFail("foo/../bar")
		ExpectPathFail("C:\\a\\b\\c.txt")
		ExpectPathFail("C:\\")
		ExpectPathFail("C:\\foo\\..\\bar")
		ExpectPathFail("C:\\..\\bar")
		ExpectPathFail("~/foo/../bar/")
		ExpectPathFail("~/../bar")
		ExpectPathFail("/foo/./")
		ExpectPathFail("/foo/.")
		ExpectPathFail("/foo/./bar")
	})
})
