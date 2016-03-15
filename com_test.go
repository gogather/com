package com

import (
	"github.com/gogather/com/log"
	. "github.com/smartystreets/goconvey/convey"
	"path/filepath"
	"testing"
)

func Test_Com(t *testing.T) {
	Convey("Test FileSystem sections", t, func() {
		Mkdir("tmp")
		CopyFile(filepath.Join("tmp", "test_copy"), "functions.go")
		So(FileExist("tmp/test_copy"), ShouldEqual, true)

		WriteFileWithCreatePath("tmp/ASD/asdfsdfa/file.log", "hello world")
		So(FileExist("tmp/ASD/asdfsdfa/file.log"), ShouldEqual, true)
	})

	Convey("Test String sections", t, func() {
		So(len(RandString(5)), ShouldEqual, 5)
		So(Strim(` hello world     `), ShouldEqual, "helloworld")
		So(HTMLEncode("请"), ShouldEqual, "&#35831;")
		So(Unicode("请"), ShouldEqual, `\u8bf7`)
	})

	Convey("Test Math sections", t, func() {
		So(Round(0.12345678, 3), ShouldEqual, 0.123)
	})

	Convey("Test Printcolor section", t, func() {
		log.Warnf("%s%d%f\n", "hello", 123, 0.456)
		log.Dangerf("%s%d%f\n", "hello", 123, 0.456)
		log.Finef("%s%d%f\n", "hello", 123, 0.456)
		log.Bluef("%s%d%f\n", "hello", 123, 0.456)
		log.Pinkf("%s%d%f\n", "hello", 123, 0.456)

		log.Warnln("hello", 123, 0.456)
		log.Dangerln("hello", 123, 0.456)
		log.Fineln("hello", 123, 0.456)
		log.Blueln("hello", 123, 0.456)
		log.Pinkln("hello", 123, 0.456)
	})
}
