package log

import (
	"bytes"
	stdlog "log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("simple logger", func() {
	var (
		newOut *bytes.Buffer
		l      Logger
	)

	BeforeEach(func() {
		newOut = &bytes.Buffer{}
		stdlog.SetOutput(newOut)
		l = NewSimple()
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

				b := newOut.Bytes()
				newOut.Reset()
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

				b := newOut.Bytes()
				newOut.Reset()
				Expect(string(b)).To(SatisfyAll(
					ContainSubstring("hi there"),
					ContainSubstring(level),
				))
			}
		})

		It("join multiple strings", func() {
			l.Debug("hi there ", "you")

			b := newOut.Bytes()
			Expect(string(b)).To(ContainSubstring("[DEBUG] hi there you"))
		})

		It("formatting", func() {
			l.Debugf("hi there %s", "you")

			b := newOut.Bytes()
			Expect(string(b)).To(ContainSubstring("[DEBUG] hi there you"))
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
					ContainSubstring("[DEBUG] hi there"),
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
					ContainSubstring("[DEBUG] hi there you"),
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
					ContainSubstring("[DEBUG] hi there you"),
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

var _ = Describe("noop logger", func() {
	var (
		l Logger
	)

	BeforeEach(func() {
		l = NewNoop()
	})

	Context("happy path", func() {
		It("does nothing", func() {
			logFuncs := map[string]func(...interface{}){
				"DEBUG": l.Debug,
				"INFO":  l.Info,
				"WARN":  l.Warn,
				"ERROR": l.Error,
			}

			for _, logFunc := range logFuncs {
				logFunc("hi there")
			}
		})

		It("does nothing on formatted", func() {
			logFuncs := map[string]func(string, ...interface{}){
				"DEBUG": l.Debugf,
				"INFO":  l.Infof,
				"WARN":  l.Warnf,
				"ERROR": l.Errorf,
			}

			for _, logFunc := range logFuncs {
				logFunc("hi %s", "there")
			}
		})

		It("join multiple strings", func() {
			l.Debug("hi there ", "you")
		})

		It("formatting", func() {
			l.Debugf("hi there %s", "you")
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
			})
		})
	})
})
