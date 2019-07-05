package numgen

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNameGenerate(t *testing.T) {
	Convey("pseduencrypt", t, func() {
		exists := make(map[uint32]bool)
		var i uint32
		for i = 0; i < 100000; i++ {
			key := pseudoEncrypt(i)
			//fmt.Println(i, key)
			_, dup := exists[key]
			So(dup, ShouldBeFalse)
			exists[key] = true
		}
	})
}

func TestNameGenerate_v2(t *testing.T) {
	Convey("pseduencrypt", t, func() {
		var i uint32
		var len uint32
		for len = 10; len > 0; len-- {
			exists := make(map[uint32]bool)
			for i = 0; i < 100000 && i < uint32(math.Pow10(int(len))); i++ {
				key := pseudoEncryptV2(i, len)
				_, dup := exists[key]
				So(dup, ShouldBeFalse)
				So(key, ShouldBeLessThan, math.Pow10(int(len+1)))
				exists[key] = true
				//	fmt.Println(i, key)
			}
		}

	})
}

func TestNameGenerator(t *testing.T) {
	Convey("generator", t, func() {
		exists := make(map[uint32]bool)
		var i uint32
		generator := GetRandNumGenerator()
		for i = 0; i < 100000; i++ {
			key := generator.Gen()
			//fmt.Println(i, key)
			_, dup := exists[key]
			So(dup, ShouldBeFalse)
			exists[key] = true
		}
	})
}
