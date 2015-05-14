package pathutils_test

import (
	. "github.com/koofr/go-pathutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NormalizeName", func() {
	It("should normalize unicode string", func() {
		Expect(NormalizeName("c")).To(Equal("c"))
		Expect(NormalizeName("s\u030c")).To(Equal("\u0161"))
		Expect(NormalizeName("a\u0000")).To(Equal("a"))
		Expect(NormalizeName("a\u001f")).To(Equal("a"))
		Expect(NormalizeName("a\u0020")).To(Equal("a "))
		Expect(NormalizeName("a\u007e")).To(Equal("a~"))
		Expect(NormalizeName("a\u007f")).To(Equal("a"))
		Expect(NormalizeName("a\u009f")).To(Equal("a"))
		Expect(NormalizeName("a\u00a0")).To(Equal("a\u00a0"))
	})
})
