# loggerZap
wrapper for zap logger

## Example
```go
package main

func main() {
	var c zlog.configLogger
	c.levCurrent = -1

	l, _ := zlog.NewExtLogger(c)
	l.Infof("info")
	l.Printf("print")
	l.Debugf("debug")
}
```
## Remaining:
### Logging with context.
### Logging with tracing.
### Inject logger in other library.
### Get current log level.
### Vary log level during ap lifecycle.
