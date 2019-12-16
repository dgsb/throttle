package throttle_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/dgsb/kipt/throttler"
)

func TestThrottler(t *testing.T) {
	tt := throttler.New(1*time.Second, 5)

	firstStart := time.Now()
	for i := 0; i < 5; i++ {
		start := time.Now()
		tt.Throttle()
		require.True(t, time.Now().Sub(start) < 10*time.Millisecond)
	}

	tt.Throttle()
	require.True(t, time.Now().Sub(firstStart) >= time.Second)
}

func TestThrottle(t *testing.T) {
	tt := throttler.New(250*time.Millisecond, 1)
	firstStart := time.Now()
	tt.Throttle()
	require.True(t, time.Now().Sub(firstStart) < 10*time.Millisecond)
	tt.Throttle()
	require.True(t, time.Now().Sub(firstStart) > 250*time.Millisecond)
	require.True(t, time.Now().Sub(firstStart) < 260*time.Millisecond)
}
