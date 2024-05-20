package pathutils_test

import (
	. "github.com/koofr/go-pathutils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validator", func() {
	It("should chech if path is valid", func() {
		Expect(IsPathValid("file.txt")).To(BeTrue())
		Expect(IsPathValid("/file.txt")).To(BeTrue())
		Expect(IsPathValid("/path\\tofile.txt")).To(BeFalse())
		Expect(IsPathValid("/path/.")).To(BeFalse())
		Expect(IsPathValid("/path/..")).To(BeFalse())
		Expect(IsPathValid("/path/../../../etc/passwd")).To(BeFalse())
		Expect(IsPathValid("..")).To(BeFalse())
		Expect(IsPathValid(".")).To(BeFalse())
	})

	It("should chech if name is valid", func() {
		Expect(IsNameValid("file.txt")).To(BeTrue())
		Expect(IsNameValid("/file.txt")).To(BeFalse())
		Expect(IsNameValid("fil\\e.txt")).To(BeFalse())
		Expect(IsNameValid(".")).To(BeFalse())
		Expect(IsNameValid("..")).To(BeFalse())
		Expect(IsNameValid("")).To(BeFalse())
	})
})
