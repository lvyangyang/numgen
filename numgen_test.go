package numgen

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNameGenerate(t *testing.T) {
	Convey("pseduencrypt", t, func() {
		exists := make(map[uint32]bool)
		var i uint32
		for i = 0; i < 100000; i++ {
			key := PseudoEncrypt(i)
			//fmt.Println(i, key)
			_, dup := exists[key]
			So(dup, ShouldBeFalse)
			exists[key] = true
		}
	})
}
