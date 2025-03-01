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
	if len(s.elements) < len(other.elements) {
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

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	iter := other

	if len(other.elements) > len(s.elements) {
		iter = s
	}

	result := NewSet[T]()
	for key := range iter.elements {
		if other.elements[key] {
			result.Add(key)
		}
	}
	return result
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for key := range s.elements {
		result.Add(key)
	}
	for key := range other.elements {
		result.Add(key)
	}
	return result
}

func (s *Set[T]) Diff(other *Set[T]) (*Set[T], *Set[T]) {
	intersection := s.Intersection(other)
	relativeComplement := s.Union(other)
	for key := range intersection.elements {
		relativeComplement.Remove(key)
	}

	return intersection, relativeComplement
}

func (s *Set[T]) ToSlice() []T {
	result := []T{}
	for key := range s.elements {
		result = append(result, key)
	}
	return result
}

func (s *Set[T]) Debug() {
	println("--------------------")
	for key := range s.elements {
		print(key)
		print(" ")
	}
	println("--------------------")

}
