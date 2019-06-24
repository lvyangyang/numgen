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
		l1 = l2
		r1 = r2
	}
	return uint32((r1 << 16) + l1)
}

func PseudoEncrypt_v2(seed uint32, dec_len uint32) uint32 {
	var bit_len uint32
	bit_len = 32
	if math.Pow10(int(dec_len)) >= float64(math.MaxUint32) {
		bit_len = 32
	} else {
		value := uint32(math.MaxUint32)
		for uint32(math.Pow10(int(dec_len))) < value>>2 {
			value >>= 2
			bit_len -= 2
		}
	}

	var rem_bit uint32 = 32 - bit_len
	var half_max uint32 = 0xffff >> (rem_bit / 2)
	var l1, r1 uint32
	l1 = (seed >> (bit_len / 2)) & uint32(half_max)
	r1 = seed & uint32(half_max)
	var l2, r2 uint32
	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ uint32(math.Round(((float64((1366*r1+150889)%714025))/714025.0)*float64(half_max)))
		l1 = l2
		r1 = r2
	}
	return uint32((l1 << (bit_len / 2)) + r1)
}
