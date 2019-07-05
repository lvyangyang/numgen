package numgen

import (
	"sync"
)

var once sync.Once

type RandNumGen struct {
	seed uint32
}

var randNumGenInstance *RandNumGen

//GetRandNumGenerator 单例生成
func GetRandNumGenerator() *RandNumGen {
	once.Do(func() {
		randNumGenInstance = &RandNumGen{0}
	})
	return randNumGenInstance
}

//Gen 生成
func (gen *RandNumGen) Gen() uint32 {
	ret := pseudoEncrypt(gen.seed)
	gen.seed++
	if ret == 0 {
		ret = pseudoEncrypt(gen.seed)
		gen.seed++
	}
	return ret
}

type RandNumGenLen struct {
	seed uint32
	len  uint32
}

var randNumGenLenInstance *RandNumGenLen

//GetRandNumLenGenerator 单例生成
func GetRandNumLenGenerator(decLen uint32) *RandNumGenLen {
	once.Do(func() {
		randNumGenLenInstance = &RandNumGenLen{seed: 0, len: decLen}
	})
	return randNumGenLenInstance
}

//Gen 生成
func (gen *RandNumGenLen) Gen() uint32 {
	ret := pseudoEncryptV2(gen.seed, gen.len)
	gen.seed++
	if ret == 0 {
		ret = pseudoEncryptV2(gen.seed, gen.len)
		gen.seed++
	}
	return ret
}
