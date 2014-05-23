package pathutils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "pathutils"
)

var _ = Describe("Path", func() {
	Describe("SplitNameExt", func() {
		It("should split name", func() {
			name, ext := SplitNameExt("file.txt")

			Expect(name).To(Equal("file"))
			Expect(ext).To(Equal("txt"))
		})

		It("should split name without ext", func() {
			name, ext := SplitNameExt("file")

			Expect(name).To(Equal("file"))
			Expect(ext).To(Equal(""))
		})

		It("should split name with multiple exts", func() {
			name, ext := SplitNameExt("file.tar.gz")

			Expect(name).To(Equal("file.tar"))
			Expect(ext).To(Equal("gz"))
		})

		It("should split name with ext only", func() {
			name, ext := SplitNameExt(".gitignore")

			Expect(name).To(Equal(""))
			Expect(ext).To(Equal("gitignore"))
		})
	})

	Describe("JoinNameExt", func() {
		It("should join name and ext", func() {
			fileName := JoinNameExt("file", "txt")

			Expect(fileName).To(Equal("file.txt"))
		})

		It("should join name without ext", func() {
			fileName := JoinNameExt("file", "")

			Expect(fileName).To(Equal("file"))
		})

		It("should join name with multiple exts", func() {
			fileName := JoinNameExt("file.tar", "gz")

			Expect(fileName).To(Equal("file.tar.gz"))
		})

		It("should join ext only", func() {
			fileName := JoinNameExt("", "gitignore")

			Expect(fileName).To(Equal(".gitignore"))
		})
	})
})
