package stream

type TokenStream[T any] interface {
	Next() *T
	Prev() *T
	Peek() *T
	EOF() bool
}

type commonTokenStream[T any] struct {
	pos    int
	tokens []T
}

func NewTokenStream[T any](input []T) TokenStream[T] {
	return &commonTokenStream[T]{tokens: input, pos: -1}
}

func (stream *commonTokenStream[T]) Next() *T {
	if stream.pos < len(stream.tokens)-1 {
		stream.pos++
		return &stream.tokens[stream.pos]
	}
	return nil
}

func (stream *commonTokenStream[T]) Prev() *T {
	if stream.pos > 0 {
		stream.pos--
		return &stream.tokens[stream.pos]
	}
	return nil
}

func (stream *commonTokenStream[T]) Peek() *T {
	if stream.pos < len(stream.tokens)-1 {
		return &stream.tokens[stream.pos+1]
	}
	return nil
}

func (stream *commonTokenStream[T]) EOF() bool {
	return stream.pos == len(stream.tokens)-1
}
