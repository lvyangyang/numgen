package numgen

import "math"

func PseudoEncrypt(seed uint32) uint32 {
	var l1, r1 uint32
	l1 = (seed >> 16) & 0xffff
	r1 = seed & 0xffff
	var l2, r2 uint32
	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ uint32(math.Round(((float64((1366*r1+150889)%714025))/714025.0)*32767))
		//r2 = l1 ^ int32(((float64((1366*r1+150889)%714025))/714025.0)*32767)
		l1 = l2
		r1 = r2
	}
	return uint32((r1 << 16) + l1)
}
