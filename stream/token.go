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

func (s *commonTokenStream[T]) Next() *T {
	if s.pos < len(s.tokens)-1 {
		s.pos++
		return &s.tokens[s.pos]
	}
	return nil
}

func (s *commonTokenStream[T]) Prev() *T {
	if s.pos > 0 {
		s.pos--
		return &s.tokens[s.pos]
	}
	return nil
}

func (s *commonTokenStream[T]) Peek() *T {
	if s.pos < len(s.tokens)-1 {
		return &s.tokens[s.pos+1]
	}
	return nil
}

func (s *commonTokenStream[T]) EOF() bool {
	return s.pos == len(s.tokens)-1
}
