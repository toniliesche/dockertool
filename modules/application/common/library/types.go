package library

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}
func (s *Stack[T]) Pop() T {
	length := len(s.items)
	item := s.items[length-1]
	s.items = s.items[:length-1]

	return item
}

func (s *Stack[T]) Items() []T {
	return s.items
}

func (s *Stack[T]) Empty() bool {
	return 0 == len(s.items)
}

func (s *Stack[T]) Update(i int, value T) {
	s.items[i] = value
}
