package com

import (
	"fmt"
	"testing"
)

func Test_RandomString(t *testing.T) {
	fmt.Println("hello: ", RandString(5))
}

func Test_Round(t *testing.T) {
	fmt.Printf("%."+"3"+"f\n", Round(0.12345678, 3))
}
