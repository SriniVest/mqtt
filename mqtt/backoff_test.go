package mqtt

import (
	"testing"
	"time"
)

func TestDurationForAttempt_NoJitter(t *testing.T) {
	b := Backoff{Base: 25, NoJitter: true}
	bInMS := time.Duration(b.Base) * time.Millisecond
	if b.DurationForAttempt(0) != bInMS {
		t.Errorf("incorrect default duration for attempt #0 (%d) = %d", b.DurationForAttempt(0)/time.Millisecond, bInMS/time.Millisecond)
	}
	var prevDuration, d time.Duration = 0, 0
	for i := 0; i < 10; i++ {
		d = b.DurationForAttempt(i)
		if !(d >= prevDuration) {
			t.Errorf("duration should be increasing between attempts. #%d (%d) > %d", i, d, prevDuration)
		}
		prevDuration = d
	}
}
