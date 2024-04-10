package base62

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEncode(t *testing.T) {
	So("0", ShouldEqual, Encode(0))
	So("0", ShouldEqual, Encode(0))
	So("1B", ShouldEqual, Encode(99))
}

func TestDecode(t *testing.T) {
	So(0, ShouldEqual, Decode("0"))
	So(99, ShouldEqual, Decode("1B"))
}

func ExampleEncode() {
	fmt.Println(Encode(99))
	// Output:
	// 1B
}

func ExampleDecode() {
	fmt.Println(Decode("1B"))
	// Output:
	// 99
}
