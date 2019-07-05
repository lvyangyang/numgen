package numgen

import (
	"sync"
	"sync/atomic"
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
	atomic.AddUint32(&gen.seed, 1)
	if ret == 0 {
		ret = pseudoEncrypt(gen.seed)
		atomic.AddUint32(&gen.seed, 1)
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

//生成
func (gen *RandNumGenLen) Gen() uint32 {
	ret := pseudoEncryptV2(gen.seed, gen.len)
	atomic.AddUint32(&gen.seed, 1)
	if ret == 0 {
		ret = pseudoEncryptV2(gen.seed, gen.len)
		atomic.AddUint32(&gen.seed, 1)
	}
	return ret
}
