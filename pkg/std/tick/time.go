package tick

const TicksPerSeconds = 60

func Seconds(ticks uint) uint {
	return ticks / TicksPerSeconds
}

func SecondsCeil(ticks uint) uint {
	if ticks%TicksPerSeconds == 0 {
		return ticks / TicksPerSeconds
	}
	return ticks/TicksPerSeconds + 1
}
