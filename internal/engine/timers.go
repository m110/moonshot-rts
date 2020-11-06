package engine

import "time"

type CountdownTimer struct {
	time      time.Duration
	resetTime time.Duration

	action func()
}

func NewCountdownTimer(duration time.Duration, action func()) *CountdownTimer {
	return &CountdownTimer{
		time:      duration,
		resetTime: duration,
		action:    action,
	}
}

func (t CountdownTimer) Done() bool {
	return t.time <= 0
}

func (t *CountdownTimer) Update(dt float64) {
	if !t.Done() {
		t.time -= time.Duration(float64(time.Second) * dt)

		if t.Done() {
			t.action()
		}
	}
}

func (t CountdownTimer) PercentDone() float64 {
	pct := float64(t.resetTime-t.time) / float64(t.resetTime)
	if pct > 1.0 {
		pct = 1.0
	}
	return pct
}
