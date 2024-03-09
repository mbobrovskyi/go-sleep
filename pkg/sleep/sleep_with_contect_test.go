package sleep_test

import (
	"context"
	"github.com/mbobrovskyi/go-sleep/pkg/sleep"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSleepWithContext(t *testing.T) {
	seconds := 1
	duration := time.Duration(seconds) * time.Second

	start := time.Now()

	sleep.WithContext(context.Background(), duration)

	assert.Equal(t, seconds, int(time.Since(start)/time.Second))
}

func TestSleepWithContextTimeout(t *testing.T) {
	seconds := 1
	duration := time.Duration(seconds) * time.Second

	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	sleep.WithContext(ctx, time.Second)

	assert.Equal(t, seconds, int(time.Since(start)/time.Second))
}
