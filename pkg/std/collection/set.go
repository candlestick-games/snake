package collection

type SetElement comparable

type Set[T SetElement] map[T]struct{}

func NewSet[T SetElement](elements ...T) Set[T] {
	s := make(Set[T], len(elements))
	for _, element := range elements {
		s[element] = struct{}{}
	}
	return s
}

func (s Set[T]) Has(key T) bool {
	_, ok := s[key]
	return ok
}

func (s Set[T]) Add(key T) {
	s[key] = struct{}{}
}

func (s Set[T]) Remove(key T) {
	delete(s, key)
}
