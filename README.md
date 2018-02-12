[![Build Status](https://travis-ci.com/InVisionApp/go-logger.svg?token=KosA43m1X3ikri8JEukQ&branch=master)](https://travis-ci.com/InVisionApp/go-logger)

# go-logger
This package provides a standard interface for logging in any go application.  
Logger interface allows you to maintain a unified interface while using a custom logger. This allows you to write log statements without dictating the specific underlying library used for logging. You can avoid vendoring of logging libraries, which is especially useful when writing shared code such as a library.  
This package also contains a simple logger and a no-op logger which both implement the interface. The simple logger is a wrapper for the standard logging library which meets this logger interface. The no-op logger can be used to easily silence all logging.  
This library is also supplemented with some additional helpers/shims for other common logging libraries such as logrus to allow them to meet the logger interface.

## Implementations

### Simple
The simple logger is a wrapper for the standard logging library which meets this logger interface. It bprovides very basic logging functionality.
```go
logger := log.NewSimple()
logger.Debug("this is a debug message")
```

### Logrus
You can choose logrus to be your logger implementation. Use the shim to meet the `ILogger` interface.
```go
// Default logrus logger
logger := log.NewLogrus(nil)
logger.Debug("this is a debug message")

// Or alternatively, you can provide your own logrus instance
myLogrus := logrus.WithField("foo", "bar")
logger := log.NewLogrus(myLogrus)
logger.Debug("this is a debug message")
```

### No-op
If you do not wish to perform any sort of logging whatsoever, you can point to a noop logger.

```go
logger := log.NewNoop()
logger.Debug("this is a debug message")
```
