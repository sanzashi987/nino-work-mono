package utils

type Set[T comparable] struct {
	elements map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]bool)}
}

func (s *Set[T]) Add(element T) {
	s.elements[element] = true
}

func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

func (s *Set[T]) Has(element T) bool {
	return s.elements[element]
}

func (s *Set[T]) IsStrictlyContains(other *Set[T]) bool {
	if len(s.elements) <= len(other.elements) {
		return false
	}

	for key := range other.elements {
		if !s.elements[key] {
			return false
		}
	}

	return true
}

func (s *Set[T]) ToRaw() map[T]bool {
	return s.elements
}
