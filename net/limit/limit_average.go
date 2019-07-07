package limit

import (
	"sync"
	"time"
)

type Limit struct {
	startTime   time.Time
	total       int64
	duration    time.Duration
	accessTotal int64
	lock        sync.Mutex
}

func New(total int64, duration time.Duration) *Limit {
	return &Limit{
		startTime:   time.Now(),
		total:       total,
		duration:    duration,
		accessTotal: 0,
	}
}

func (l *Limit) Reset(t time.Time) {
	l.startTime = t
	l.accessTotal = 0
}

func (l *Limit) isLimit() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.accessTotal == l.total-1 {
		t := time.Now()
		if t.Sub(l.startTime) >= l.duration {
			l.Reset(t)
			return true
		} else {
			return false
		}
	} else {
		l.accessTotal += 1
		return true
	}
}

func main() {

}
