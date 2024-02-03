package tick

type Timer struct {
	ticker *Ticker

	start    uint
	duration uint
	stopped  bool
}

func (t *Timer) IsStarted() bool {
	return t != nil && !t.stopped
}

func (t *Timer) Reset(duration uint) {
	t.start = t.ticker.ticks
	t.duration = duration
	t.stopped = false
}

func (t *Timer) Stop() {
	t.stopped = true
}

func (t *Timer) Wait() bool {
	if t == nil || t.stopped {
		return false
	}

	if t.ticker.ticks >= t.start+t.duration {
		t.stopped = true
		return true
	}

	return false
}
