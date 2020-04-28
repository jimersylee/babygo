// runtime
package runtime

import "unsafe"

var heap [4096]uint8

var heapHead uintptr
var heapCurrent uintptr
var heapTail uintptr

func heapInit() {
	heapHead = uintptr(unsafe.Pointer(&heap[0])) // brk(0)
	heapTail = heapHead + 4096                   // brk(heapHead + heapSize)
	heapCurrent = heapHead
}

func memzeropad(addr uintptr, size uintptr) {
	var p *uint8 = (*uint8)(unsafe.Pointer(addr))
	var isize int = int(size)
	var i int
	var up uintptr
	for i = 0; i < isize; i++ {
		*p = 0
		up = uintptr(unsafe.Pointer(p)) + 1
		p = (*uint8)(unsafe.Pointer(up))
	}
}

func memcopy(src uintptr, dst uintptr, length int) {
	var i int
	var srcp *uint8
	var dstp *uint8
	for i = 0; i < length; i++ {
		srcp = (*uint8)(unsafe.Pointer(src + uintptr(i)))
		dstp = (*uint8)(unsafe.Pointer(dst + uintptr(i)))
		*dstp = *srcp
	}
}

func malloc(size uintptr) uintptr {
	if heapCurrent+size > heapTail {
		panic("malloc exceeds heap capacity")
		return 0
	}
	var r uintptr
	r = heapCurrent
	heapCurrent = heapCurrent + size
	memzeropad(r, size)
	return r
}

func makeSlice(elmSize int, slen int, scap int) (uintptr, int, int) {
	var size uintptr = uintptr(elmSize * scap)
	var addr uintptr = malloc(size)
	return addr, slen, scap
}

// Actually this is an alias to makeSlice
func makeSlice1(elmSize int, slen int, scap int) []uint8
func makeSlice8(elmSize int, slen int, scap int) []int
func makeSlice16(elmSize int, slen int, scap int) []string
func makeSlice24(elmSize int, slen int, scap int) [][]int

func append1(old []uint8, elm uint8) (uintptr, int, int) {
	var new_ []uint8
	var elmSize int = 1

	var oldlen int = len(old)
	var newlen int = oldlen + 1

	if cap(old) >= newlen {
		new_ = old[0:newlen]
	} else {
		var newcap int
		if oldlen == 0 {
			newcap = 1
		} else {
			newcap = oldlen * 2
		}
		new_ = makeSlice1(elmSize, newlen, newcap)
		var oldSize int = oldlen * elmSize
		if oldlen > 0 {
			memcopy(uintptr(unsafe.Pointer(&old[0])), uintptr(unsafe.Pointer(&new_[0])), oldSize)
		}
	}

	new_[oldlen] = elm
	return uintptr(unsafe.Pointer(&new_[0])), newlen, cap(new_)
}

func append8(old []int, elm int) (uintptr, int, int) {
	var new_ []int
	var elmSize int = 8

	var oldlen int = len(old)
	var newlen int = oldlen + 1

	if cap(old) >= newlen {
		new_ = old[0:newlen]
	} else {
		var newcap int
		if oldlen == 0 {
			newcap = 1
		} else {
			newcap = oldlen * 2
		}
		new_ = makeSlice8(elmSize, newlen, newcap)
		var oldSize int = oldlen * elmSize
		if oldlen > 0 {
			memcopy(uintptr(unsafe.Pointer(&old[0])), uintptr(unsafe.Pointer(&new_[0])), oldSize)
		}
	}

	new_[oldlen] = elm
	return uintptr(unsafe.Pointer(&new_[0])), newlen, cap(new_)
}

func append16(old []string, elm string) (uintptr, int, int) {
	var new_ []string
	var elmSize int = 16

	var oldlen int = len(old)
	var newlen int = oldlen + 1

	if cap(old) >= newlen {
		new_ = old[0:newlen]
	} else {
		var newcap int
		if oldlen == 0 {
			newcap = 1
		} else {
			newcap = oldlen * 2
		}
		new_ = makeSlice16(elmSize, newlen, newcap)
		var oldSize int = oldlen * elmSize
		if oldlen > 0 {
			memcopy(uintptr(unsafe.Pointer(&old[0])), uintptr(unsafe.Pointer(&new_[0])), oldSize)
		}
	}

	new_[oldlen] = elm
	return uintptr(unsafe.Pointer(&new_[0])), newlen, cap(new_)
}

func append24(old [][]int, elm []int) (uintptr, int, int) {
	var new_ [][]int
	var elmSize int = 24

	var oldlen int = len(old)
	var newlen int = oldlen + 1

	if cap(old) >= newlen {
		new_ = old[0:newlen]
	} else {
		var newcap int
		if oldlen == 0 {
			newcap = 1
		} else {
			newcap = oldlen * 2
		}
		new_ = makeSlice24(elmSize, newlen, newcap)
		var oldSize int = oldlen * elmSize
		if oldlen > 0 {
			memcopy(uintptr(unsafe.Pointer(&old[0])), uintptr(unsafe.Pointer(&new_[0])), oldSize)
		}
	}

	new_[oldlen] = elm
	return uintptr(unsafe.Pointer(&new_[0])), newlen, cap(new_)
}

func panic(s string) {
	print(s)
	//exit(1)
}

func nop() {
}

func nop1() {
}

func nop2() {
}

func nop3() {
}

func catstrings(a string, b string) string {
	var totallen int
	var r []uint8
	totallen = len(a) + len(b)
	r = make([]uint8, totallen, totallen)
	var i int
	for i = 0; i < len(a); i = i + 1 {
		r[i] = a[i]
	}
	var j int
	for j = 0; j < len(b); j = j + 1 {
		r[i+j] = b[j]
	}
	return string(r)
}

func cmpstrings(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var i int
	for i = 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
