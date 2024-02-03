package call

import "runtime"

const pcCap = 32

type Pos = [pcCap]uintptr

func GetPos() Pos {
	pc := Pos{}
	n := runtime.Callers(1, pc[:])
	if n >= pcCap {
		panic("pc overflow")
	}
	return pc
}
