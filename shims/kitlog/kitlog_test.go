package kitlog

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"

	"github.com/InVisionApp/go-logger"
	kitlog "github.com/go-kit/kit/log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"os"
)

var _ = Describe("satisfies interface", func() {
	var _ log.Logger = &shim{}
})

var _ = Describe("kitlog logger", func() {
	var (
		newOut *bytes.Buffer
		l      log.Logger
	)
	BeforeEach(func() {
		newOut = &bytes.Buffer{}
		l = New(kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(newOut)))
	})
	Context("spaceSep", func() {
		It("works on even length slices", func() {
			s := []interface{}{
				1, "cat", errors.New("foo"), struct{ Name string }{"bar"},
			}
			s = spaceSep(s)
			Expect(len(s)).To(Equal(7))

			sstr := fmt.Sprint(s...)
			Expect(sstr).To(Equal("1 cat foo {bar}"))
		})
		It("works on odd length slices", func() {
			s := []interface{}{
				1, "cat", errors.New("foo"), struct{ Name string }{"bar"}, []string{},
			}
			s = spaceSep(s)
			Expect(len(s)).To(Equal(9))

			sstr := fmt.Sprint(s...)
			Expect(sstr).To(Equal("1 cat foo {bar} []"))
		})
		It("works on big slices", func() {
			s := make([]interface{}, 10000)
			for i := 0; i < 10000; i++ {
				s[i] = rand.Intn(10)
			}
			s = spaceSep(s)
			Expect(len(s)).To(Equal(19999))

			sstr := fmt.Sprint(s...)
			Expect(len(sstr)).To(Equal(19999))
		})
		It("works on little slices", func() {
			s := []interface{}{1, 2}

			s = spaceSep(s)
			Expect(len(s)).To(Equal(3))

			sstr := fmt.Sprint(s...)
			Expect(sstr).To(Equal("1 2"))
		})
		It("works on really little slices", func() {
			s := []interface{}{1}

			s = spaceSep(s)
			Expect(len(s)).To(Equal(1))

			sstr := fmt.Sprint(s...)
			Expect(sstr).To(Equal("1"))
		})
	})
	Context("log funcs", func() {
		It("prints all the log levels", func() {
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
		It("prints all the log levels: *ln", func() {
			logFuncs := map[string]func(...interface{}){
				"debug?": l.Debugln,
				"info":   l.Infoln,
				"warn":   l.Warnln,
				"error?": l.Errorln,
			}
			for level, logFunc := range logFuncs {
				logFunc("hi", "there")

				b := newOut.Bytes()
				newOut.Reset()
				Expect(string(b)).To(SatisfyAll(
					ContainSubstring("hi there"),
					MatchRegexp(level),
				))
			}
		})
		It("prints all the log levels: *f", func() {
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
		It("nil logger", func() {
			// need to intercept stdout
			// https://stackoverflow.com/a/10476304
			old := os.Stdout

			r, w, _ := os.Pipe()
			os.Stdout = w

			l = New(nil)
			l.Debug("i am default")

			outC := make(chan string)
			go func() {
				var buf bytes.Buffer
				io.Copy(&buf, r)
				outC <- buf.String()
			}()
			w.Close()
			os.Stdout = old

			out := <-outC
			Expect(out).To(MatchRegexp(`level=debug msg="i am default"`))
		})
	})
	Context("fields", func() {
		It("", func() {
			l = l.WithFields(log.Fields{
				"foo": "bar",
				"tf":  true,
				"pet": "cat",
				"age": 1,
			})

			l.Debug("hi there")
			b := newOut.Bytes()

			Expect(string(b)).To(SatisfyAll(
				ContainSubstring("hi there"),
				ContainSubstring("level=debug"),
				ContainSubstring("foo=bar"),
				ContainSubstring("tf=true"),
				ContainSubstring("pet=cat"),
				ContainSubstring("age=1"),
			))
		})
	})
})
