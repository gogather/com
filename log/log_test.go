package log

import (
	"testing"
)

func Test_PrintColor(t *testing.T) {
	var err error

	_, err = Warnf("%s%d%f\n", "hello", 123, 0.456)
	_, err = Dangerf("%s%d%f\n", "hello", 123, 0.456)
	_, err = Finef("%s%d%f\n", "hello", 123, 0.456)
	_, err = Bluef("%s%d%f\n", "hello", 123, 0.456)
	_, err = Pinkf("%s%d%f\n", "hello", 123, 0.456)

	_, err = Warnln("hello", 123, 0.456)
	_, err = Dangerln("hello", 123, 0.456)
	_, err = Fineln("hello", 123, 0.456)
	_, err = Blueln("hello", 123, 0.456)
	_, err = Pinkln("hello", 123, 0.456)

	if err != nil {
		t.Error(err)
	}
}
