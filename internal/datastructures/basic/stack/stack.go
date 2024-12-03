package stack

type Stack[T any] struct {
	Data []T
}

func (s *Stack[T]) Push(value T) {
	s.Data = append(s.Data, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		return *new(T), false
	}

	length := len(s.Data)
	value := s.Data[length-1]
	s.Data = s.Data[:length-1]

	return value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		return *new(T), false
	}

	return s.Data[len(s.Data)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Size() int {
	return len(s.Data)
}
