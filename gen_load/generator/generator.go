package generator

import (
	"math/rand"
	"time"
)

// Generator generates messages, limit rate in bytes per second, with per-second granularity.
type Generator struct {
	RateBytesPerSec   int64
	rateLimitedSecond int64
	rateUsed          int64
}

// NewGenerator creates a new  generator.
func NewGenerator(rateBytesPerSec int64) *Generator {
	return &Generator{rateBytesPerSec, 0, 0}
}

// Read reads new message from generator or returns a RateLimitError if ratelimited.
func (mg *Generator) Read(p []byte) (n int, err error) {
	currentSecond := time.Now().Unix()
	if currentSecond > mg.rateLimitedSecond {
		mg.rateLimitedSecond = currentSecond
		mg.rateUsed = 0
	}

	if mg.rateUsed >= mg.RateBytesPerSec {
		// TODO: Move to debug output
		// fmt.Printf("SEC %d: rate limited\n", currentSecond)
		ts := time.Now().UnixNano()
		limitTTL := 1e9 - ts%1e9
		err = &RateLimitedError{"Rate limited", time.Duration(limitTTL) * time.Nanosecond}
		return 0, err
	}

	// TODO: Migrate from serializer to random string generator, avoid string/[]byte conversions
	n, err = rand.Read(p)
	mg.rateUsed += int64(len(p))
	// TODO: Move to debug output
	// fmt.Printf("SEC %d: Got %d bytes, used %d of %d bytes\n", currentSecond, len(p), mg.rateUsed, mg.RateBytesPerSec)
	return n, err
}

// RateLimitedError specifies a situation when read was rate limited
type RateLimitedError struct {
	err      string
	LimitTTL time.Duration
}

func (e *RateLimitedError) Error() string {
	return e.err
}
