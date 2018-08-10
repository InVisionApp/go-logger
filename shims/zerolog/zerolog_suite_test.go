package zerolog_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestZerolog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Zerolog Suite")
}
