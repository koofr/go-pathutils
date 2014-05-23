package pathutils_test

import (
	. "github.com/koofr/go-pathutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func check(path string, existing []string, newPath string, retries int, expectErr bool) {
	existsFn := func(path string) bool {
		for _, e := range existing {
			if path == e {
				return true
			}
		}
		return false
	}

	path, err := UnusedFilename(existsFn, path, retries)

	if expectErr {
		Expect(err).To(HaveOccurred())
	} else {
		Expect(err).NotTo(HaveOccurred())
		Expect(path).To(Equal(newPath))
	}
}

var _ = Describe("UnusedFilename", func() {
	It("should return same path if it doesn not exist", func() {
		check("/path/to/file.txt", []string{}, "/path/to/file.txt", 10, false)
	})

	It("should incremented path if it already exists", func() {
		check("file.txt", []string{"file.txt"}, "file (1).txt", 10, false)
	})

	It("should not return error if retries do not exceed limit", func() {
		check("file.txt", []string{"file.txt", "file (1).txt"}, "file (2).txt", 3, false)
	})

	It("should return error if too many retries", func() {
		check("file.txt", []string{"file.txt", "file (1).txt"}, "", 2, true)
	})
})
