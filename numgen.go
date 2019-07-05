package numgen

import (
	"math"
)

func pseudoEncrypt(seed uint32) uint32 {
	var l1, r1 uint32
	l1 = (seed >> 16) & 0xffff
	r1 = seed & 0xffff
	var l2, r2 uint32
	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ uint32(math.Round(((float64((1366*r1+150889)%714025))/714025.0)*32767))
		l1 = l2
		r1 = r2
	}
	return uint32((r1 << 16) + l1)
}

func pseudoEncryptV2(seed uint32, decLen uint32) uint32 {
	var bitLen uint32
	bitLen = 32
	if math.Pow10(int(decLen)) >= float64(math.MaxUint32) {
		bitLen = 32
	} else {
		value := uint32(math.MaxUint32)
		for uint32(math.Pow10(int(decLen))) < value>>2 {
			value >>= 2
			bitLen -= 2
		}
	}

	var remBit = 32 - bitLen
	var halfMax uint32 = 0xffff >> (remBit / 2)
	var l1, r1 uint32
	l1 = (seed >> (bitLen / 2)) & uint32(halfMax)
	r1 = seed & uint32(halfMax)
	var l2, r2 uint32
	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ uint32(math.Round(((float64((1366*r1+150889)%714025))/714025.0)*float64(halfMax)))
		l1 = l2
		r1 = r2
	}
	return uint32((l1 << (bitLen / 2)) + r1)
}
