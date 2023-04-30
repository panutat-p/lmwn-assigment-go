package src

type Set map[string]struct{}

func NewSet() Set {
	return make(map[string]struct{})
}

func NewSetFromSlice(sl []string) Set {
	var (
		m = make(map[string]struct{})
	)
	for _, e := range sl {
		m[e] = struct{}{}
	}
	return m
}

func (s *Set) Add(key string) {
	(*s)[key] = struct{}{}
}

func (s *Set) Has(key string) bool {
	_, ok := (*s)[key]
	return ok
}
