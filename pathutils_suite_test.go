package pathutils_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPathutils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pathutils Suite")
}
