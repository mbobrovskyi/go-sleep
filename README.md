# Go sleep

This is a tool that can help you set timeouts with context

### How to use:

```go
package main

import (
	"context"
	"github.com/mbobrovskyi/go-sleep/pkg/sleep"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sleep.WithContext(ctx, time.Second)
}

```