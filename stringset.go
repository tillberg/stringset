package stringset

type StringSet struct {
	strMap map[string]bool
}

func New(elements ...string) *StringSet {
	s := &StringSet{}
	s.Clear()
	for _, el := range elements {
		s.Add(el)
	}
	return s
}
func (s *StringSet) Len() int {
	return len(s.strMap)
}
func (s *StringSet) Add(str string) bool {
	if _, exists := s.strMap[str]; exists {
		return false
	}
	s.strMap[str] = true
	return true
}
func (s *StringSet) Remove(str string) {
	delete(s.strMap, str)
}
func (s *StringSet) Has(str string) bool {
	_, has := s.strMap[str]
	return has
}
func (s *StringSet) All() []string {
	l := make([]string, 0, len(s.strMap))
	for str, _ := range s.strMap {
		l = append(l, str)
	}
	return l
}
func (s *StringSet) Clear() {
	s.strMap = make(map[string]bool)
}
func (s *StringSet) Raw() map[string]bool {
	return s.strMap
}
func (s *StringSet) Equal(other *StringSet) bool {
	if len(s.strMap) != len(other.strMap) {
		return false
	}
	for k := range s.strMap {
		if !other.strMap[k] {
			return false
		}
	}
	return true
}
