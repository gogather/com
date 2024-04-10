package stream

import (
	"fmt"
	"testing"
)

func TestTokenStream(t *testing.T) {
	input := "x = 1 + 2 * 3"
	stream := NewTokenStream([]byte(input))

	for !stream.EOF() {
		token := stream.Next()
		fmt.Printf("Token -> %c\n", *token)
	}

	fmt.Println("stream:", stream)
}
