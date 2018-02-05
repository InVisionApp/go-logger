package logrus

import (
	"bytes"

	"github.com/InVisionApp/go-logger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

var _ = Describe("logrus logger", func() {
	var (
		newOut *bytes.Buffer
		l      log.Logger
	)

	BeforeEach(func() {
		newOut = &bytes.Buffer{}
		logrus.SetOutput(newOut)
		logrus.SetLevel(logrus.DebugLevel)
		l = NewLogrus(nil)
	})

	Context("happy path", func() {
		It("prints all log levels", func() {
			logFuncs := map[string]func(...interface{}){
				"debug": l.Debug,
				"info":  l.Info,
				"warn":  l.Warn,
				"error": l.Error,
			}

			for level, logFunc := range logFuncs {
				logFunc("hi there")

				b := newOut.Bytes()
				newOut.Reset()
				Expect(string(b)).To(SatisfyAll(
					ContainSubstring("hi there"),
					ContainSubstring("level="+level),
				))
			}
		})

		It("prints all log levels on formatted", func() {
			logFuncs := map[string]func(string, ...interface{}){
				"debug": l.Debugf,
				"info":  l.Infof,
				"warn":  l.Warnf,
				"error": l.Errorf,
			}

			for level, logFunc := range logFuncs {
				logFunc("hi %s", "there")

				b := newOut.Bytes()
				newOut.Reset()
				Expect(string(b)).To(SatisfyAll(
					ContainSubstring("hi there"),
					ContainSubstring("level="+level),
				))
			}
		})

		It("join multiple strings", func() {
			l.Debug("hi there ", "you")

			b := newOut.Bytes()
			Expect(string(b)).To(SatisfyAll(
				ContainSubstring("hi there you"),
				ContainSubstring("level=debug"),
			))
		})

		It("formatting", func() {
			l.Debugf("hi there %s", "you")

			b := newOut.Bytes()
			Expect(string(b)).To(SatisfyAll(
				ContainSubstring("hi there you"),
				ContainSubstring("level=debug"),
			))
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

				b := newOut.Bytes()
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

				b := newOut.Bytes()
				Expect(string(b)).To(SatisfyAll(
					ContainSubstring("hi there you"),
					ContainSubstring("foo=bar"),
					ContainSubstring("baz=2"),
				))

				newOut.Reset()

				// should not see any of the other fields

				l.WithFields(map[string]interface{}{
					"biz": "bar",
					"buz": 2,
				}).Debugf("hi there %s", "you")

				bb := newOut.Bytes()
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
