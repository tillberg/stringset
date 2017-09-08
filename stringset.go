package stringset

type StringSet struct {
	strMap map[string]struct{}
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
	s.strMap[str] = struct{}{}
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
	s.strMap = make(map[string]struct{})
}
func (s *StringSet) Raw() map[string]struct{} {
	return s.strMap
}
func (s *StringSet) Equal(other *StringSet) bool {
	if len(s.strMap) != len(other.strMap) {
		return false
	}
	for k := range s.strMap {
		if _, exists := other.strMap[k]; !exists {
			return false
		}
	}
	return true
}

func (s *StringSet) Difference(other *StringSet) *StringSet {
	diff := &StringSet{}
	diff.Clear()
	for k := range s.strMap {
		if _, exists := other.strMap[k]; !exists {
			diff.Add(k)
		}
	}
	return diff
}

func (s *StringSet) Intersection(other *StringSet) *StringSet {
	inter := &StringSet{}
	inter.Clear()
	for k := range s.strMap {
		if _, exists := other.strMap[k]; exists {
			inter.Add(k)
		}
	}
	return inter
}

func (s *StringSet) Union(other *StringSet) *StringSet {
	union := &StringSet{}
	union.Clear()
	for k := range s.strMap {
		union.Add(k)
	}
	for k := range other.strMap {
		union.Add(k)
	}
	return union
}