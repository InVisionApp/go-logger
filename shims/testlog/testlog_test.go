package testlog

import (
	"github.com/InVisionApp/go-logger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("meets the interface", func() {
	var _ log.Logger = &TestLogger{}
})

var _ = Describe("test logger", func() {
	var (
		testOut *TestLogger
		l       log.Logger
	)

	BeforeEach(func() {
		testOut = New()
		l = testOut
	})

	Context("happy path", func() {
		It("prints all log levels", func() {
			logFuncs := map[string]func(...interface{}){
				"DEBUG": l.Debug,
				"INFO":  l.Info,
				"WARN":  l.Warn,
				"ERROR": l.Error,
			}

			for level, logFunc := range logFuncs {
				logFunc("hi there")

				b := testOut.Bytes()
				testOut.Reset()
				Expect(string(b)).To(SatisfyAll(
					ContainSubstring("hi there"),
					ContainSubstring(level),
				))
			}
		})

		It("prints all log levels on formatted", func() {
			logFuncs := map[string]func(string, ...interface{}){
				"DEBUG": l.Debugf,
				"INFO":  l.Infof,
				"WARN":  l.Warnf,
				"ERROR": l.Errorf,
			}

			for level, logFunc := range logFuncs {
				logFunc("hi %s", "there")

				b := testOut.Bytes()
				testOut.Reset()
				Expect(string(b)).To(SatisfyAll(
					ContainSubstring("hi there"),
					ContainSubstring(level),
				))
			}
		})

		It("join multiple strings", func() {
			l.Debug("hi there ", "you")

			b := testOut.Bytes()
			Expect(string(b)).To(SatisfyAll(
				ContainSubstring("hi there you"),
				ContainSubstring("DEBUG"),
			))
		})

		It("formatting", func() {
			l.Debugf("hi there %s", "you")

			b := testOut.Bytes()
			Expect(string(b)).To(SatisfyAll(
				ContainSubstring("hi there you"),
				ContainSubstring("DEBUG"),
			))
		})

		Context("Bytes", func() {
			It("returns the bytes in the buffer", func() {
				testOut.buf.WriteString("foo")

				b := testOut.Bytes()

				Expect(string(b)).To(Equal("foo"))
			})
		})

		Context("CallCount", func() {
			It("returns the call count", func() {
				l.Info("foo")
				l.Info("foo")
				l.Info("foo")

				Expect(testOut.CallCount()).To(Equal(3))
			})
		})

		Context("reset", func() {
			It("resets the buffer and the call count", func() {
				l.Info("foo")

				testOut.Reset()

				Expect(testOut.Bytes()).To(BeEmpty())
				Expect(testOut.CallCount()).To(Equal(0))
			})
		})

		Context("with fields", func() {
			It("appends to preexisting fields", func() {
				withFields := l.WithFields(map[string]interface{}{
					"foo": "oldval",
					"baz": "origval",
				})

				withFields.WithFields(map[string]interface{}{
					"foo": "newval",
					"biz": "buzz",
				}).Debug("hi there")

				b := testOut.Bytes()
				Expect(string(b)).To(SatisfyAll(
					ContainSubstring("hi there"),
					ContainSubstring("foo=newval"),
					ContainSubstring("baz=origval"),
					ContainSubstring("biz=buzz"),
				))

			})

			It("creates a copy", func() {
				l.WithFields(map[string]interface{}{
					"foo": "bar",
					"baz": 2,
				}).Debug("hi there ", "you")

				b := testOut.Bytes()
				Expect(string(b)).To(SatisfyAll(
					ContainSubstring("hi there you"),
					ContainSubstring("foo=bar"),
					ContainSubstring("baz=2"),
				))

				testOut.Reset()

				// should not see any of the other fields

				l.WithFields(map[string]interface{}{
					"biz": "bar",
					"buz": 2,
				}).Debugf("hi there %s", "you")

				bb := testOut.Bytes()
				Expect(string(bb)).To(SatisfyAll(
					ContainSubstring("hi there you"),
					ContainSubstring("biz=bar"),
					ContainSubstring("buz=2"),
				))
				Expect(string(bb)).ToNot(SatisfyAll(
					ContainSubstring("foo=bar"),
					ContainSubstring("baz=2"),
				))
			})
		})
	})
})
