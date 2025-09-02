# errutil

Go utilities for consistent, predictable error handling — with logging, wrapping, and retry with backoff.

## Why use it?
- Standardizes how you log and propagate errors
- Simple retry helpers with backoff
- Plays nicely with your logger (pairs great with `logutil`)

## Installation
```sh
go get github.com/alice-bnuy/errutil
```

## Quick start
The example below initializes a logger, sets up the global error handler, and runs an operation with retries:

```go
package main

import (
    "context"
    "errors"
    "fmt"
    "time"

    "github.com/alice-bnuy/errutil"
    "github.com/alice-bnuy/logutil"
)

func main() {
    // 1) Initialize a logger (console enabled, file optional)
    logger, _ := logutil.NewLogger(logutil.LoggerConfig{
        Level:         logutil.InfoLevel,
        LogDir:        logutil.LogsDirPath,
        EnableConsole: true,
        EnableFile:    false,
        IncludeCaller: true,
    })
    defer logger.Close()

    // 2) Register errutil's global error handler
    _ = errutil.InitializeGlobalErrorHandler(logger)

    // 3) Run an operation with retry (3 attempts)
    err := errutil.RetryOperation(context.Background(), "fetch-data", 3, func() error {
        // Simulate a flaky operation: fail sometimes
        if time.Now().UnixNano()%3 != 0 {
            return errors.New("transient network error")
        }
        return nil // success!
    })

    if err != nil {
        fmt.Println("Operation failed:", err)
        return
    }
    fmt.Println("Operation completed successfully")
}
```

## Other handy helpers
- `HandleValidationError` and `HandleConfigError` for consistent messages and context
- `LogAndWrapError` to attach context while preserving error information
- `EnsureSuccess` for critical operations (fails the process on error)

## License
AGPL-3.0
