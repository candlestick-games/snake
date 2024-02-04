package call

import "runtime"

const pcCap = 32

type Pos = [pcCap]uintptr // TODO: Check if calculation hash is better

func GetPos(ids ...uintptr) Pos {
	pc := Pos{}
	n := runtime.Callers(1, pc[:])
	m := len(ids)
	if n+m > pcCap {
		panic("pc overflow")
	}
	if copy(pc[n:], ids) != m {
		panic("pc ids overflow")
	}
	return pc
}
