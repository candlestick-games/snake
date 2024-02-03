package tick

import (
	"github.com/candlestick-games/snake/pkg/std/call"
)

type Ticker struct {
	ticks   uint
	waiters map[call.Pos]uint
}

func NewTicker() *Ticker {
	return &Ticker{
		ticks:   0,
		waiters: make(map[call.Pos]uint),
	}
}

func (t *Ticker) Update() {
	t.ticks++
}

func (t *Ticker) Every(ticks uint) bool {
	return t.ticks%ticks == 0
}

func (t *Ticker) NewTimer() *Timer {
	return &Timer{
		ticker:   t,
		start:    0,
		duration: 0,
		stopped:  true,
	}
}

func (t *Ticker) StartTimer(duration uint) *Timer {
	return &Timer{
		ticker:   t,
		start:    t.ticks,
		duration: duration,
		stopped:  false,
	}
}

func (t *Ticker) Wait(ticks uint) bool {
	pos := call.GetPos()

	start, ok := t.waiters[pos]
	if !ok {
		start = t.ticks
		t.waiters[pos] = start
	}

	if t.ticks >= start+ticks {
		delete(t.waiters, pos)
		return true
	}

	return false
}
