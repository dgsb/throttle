// package throttle implements throttling on Runner interface.
package throttle

import (
	"container/ring"
	"time"
)

type Throttler struct {
	period       time.Duration
	numOccurence uint
	r            *ring.Ring
}

func New(period time.Duration, numOccurence uint) *Throttler {
	if numOccurence == 0 {
		return nil
	}
	r := ring.New(int(numOccurence))
	t := &Throttler{
		period:       period,
		numOccurence: numOccurence,
		r:            r,
	}
	return t
}

func (t *Throttler) needThrottle() bool {
	if t.r.Next().Value == nil {
		return false
	}
	if time.Now().Sub(*t.r.Next().Value.(*time.Time)) > t.period {
		return false
	}
	return true
}

func (t *Throttler) Throttle() {
	if t.needThrottle() {
		time.Sleep(
			t.r.Next().Value.(*time.Time).Add(t.period).Sub(time.Now()))
	}
	t.r = t.r.Next()
	t.UpdateTimestamp()
}

func (t *Throttler) UpdateTimestamp() {
	now := time.Now()
	t.r.Value = &now
}
