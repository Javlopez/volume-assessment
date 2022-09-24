package calculate

type Set map[string]bool

func NewSet() *Set {
	return &Set{}
}

func (s Set) Has(node string) bool {
	if _, ok := s[node]; ok {
		return true
	}
	return false
}

func (s Set) Add(node string) {
	s[node] = true
}
