package testlog

import (
	"testing"

	"github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSuiteNameSuite(t *testing.T) {
	// reduce the noise when testing
	logrus.SetLevel(logrus.FatalLevel)

	RegisterFailHandler(Fail)
	RunSpecs(t, "SuiteName Suite")
}
