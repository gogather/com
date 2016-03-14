package com

import (
	"fmt"
	"path/filepath"
	"testing"
)

func Test_RandomString(t *testing.T) {
	fmt.Println("hello: ", RandString(5))
}

func Test_Round(t *testing.T) {
	fmt.Printf("%."+"3"+"f\n", Round(0.12345678, 3))
}

func Test_Strim(t *testing.T) {
	str := ` hello world
	hi
	`

	if "helloworldhi" == Strim(str) {
		fmt.Println("Strim Passed")
	} else {
		t.Error("failed in Strim")
	}
}

func Test_Copy(t *testing.T) {
	Mkdir("tmp")
	CopyFile(filepath.Join("tmp", "test_copy"), "functions.go")
}

func Test_encode(t *testing.T) {
	str := HTMLEncode("请")
	fmt.Println("[HTMLEncode]", str)

	str = Unicode("请")
	fmt.Println("[Unicode]", str)
}

func Test_Mkdir(t *testing.T) {
	err := WriteFileWithCreatePath("tmp/ASD/asdfsdfa/file.log", "hello world")
	fmt.Println(err)
}
