[![codecov](https://codecov.io/gh/dangkaka/go-simple-logger/branch/master/graph/badge.svg)](https://codecov.io/gh/dangkaka/go-simple-logger)  [![Build Status](https://travis-ci.org/dangkaka/go-simple-logger.svg?branch=master)](https://travis-ci.org/dangkaka/go-simple-logger)

# logger

A simple a logger with json format

## Usage
```go
package main
import "github.com/dangkaka/logger"

func main() {
    //LevelFatal   = 0
    //LevelError   = 1
    //LevelWarning = 2
    //LevelInfo    = 3
    //LevelDebug   = 4
    log.SetLevel(3)
    err := "xxx"
    log.Error("error: ", err)
}
```

*Result:*
```
{"level":"ERROR","msg":"error:  xxx ","time":"Friday, 29-Jun-18 23:17:32 +07","track":" [/Users/user/go/src/github.com/dangkaka/test.go:12] [/usr/local/Cellar/go/1.10.1/libexec/src/runtime/proc.go:198] [/usr/local/Cellar/go/1.10.1/libexec/src/runtime/asm_amd64.s:2361]"}
```
