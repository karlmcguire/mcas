package mcas

import _ "unsafe"

func PID() int {
	pid := procPin()
	procUnpin()
	return pid
}

//go:linkname procPin runtime.procPin
//go:nosplit
func procPin() int

//go:linkname procUnpin runtime.procUnpin
//go:nosplit
func procUnpin()
