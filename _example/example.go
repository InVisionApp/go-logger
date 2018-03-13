package main

import (
	"fmt"
	"time"

	"github.com/InVisionApp/go-logger"
	"github.com/InVisionApp/go-logger/shims/_fake"
	"github.com/InVisionApp/go-logger/shims/logrus"
	"github.com/InVisionApp/go-logger/shims/testlog"
)

func main() {
	/**********************
	 Simple, no-op, logrus
	**********************/

	loggers := map[string]log.Logger{
		"Simple": log.NewSimple(),
		"No-op":  log.NewNoop(),
		"Logrus": logrus.New(nil),
	}

	// sleeps for print order
	for name, lg := range loggers {
		time.Sleep(time.Millisecond * 10)
		fmt.Printf("__%s Logger__________\n", name)
		time.Sleep(time.Millisecond * 10)

		lg.Info(name, "logger")
		lg.WithFields(log.Fields{"foo": "bar"}).Errorf("%s logger", name)

		time.Sleep(time.Millisecond * 10)
		fmt.Println()
	}

	time.Sleep(time.Millisecond * 10)

	/************
	 Test Logger
	************/

	fmt.Println("__Test Logger__________")
	tl := testlog.New()
	tl.Infof("this is a test %s", "log")
	fmt.Print(string(tl.Bytes()))
	fmt.Println("call count:", tl.CallCount())
	fmt.Println()

	/************
	  Fake Logger
	 ************/

	fmt.Println("__Fake Logger__________")
	fakeLog := &fake.FakeLogger{}

	fakeLog.Debug("this is a ", "fakelog")
	fmt.Println("call count:", fakeLog.DebugCallCount())
}
