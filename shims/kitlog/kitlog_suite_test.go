package kitlog_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestKitlog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kitlog Suite")
}
