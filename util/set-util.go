package util

// Set 集合实现

//var Empty struct{} //空结构体变量的内存占用大小为0
//作为常量填充到map的value
var Empty = struct{}{}

type Set struct {
	m map[interface{}]struct{}
}

func NewSet(items ...interface{}) *Set {
	s := &Set{
		m: make(map[interface{}]struct{}),
	}
	for _, item := range items {
		if values, ok := item.([]int); ok {
			for _, value := range values {
				s.m[value] = Empty
			}
		} else {
			s.m[items] = Empty
		}
	}
	return s
}

func (s *Set) Add(items ...interface{}) {
	for _, item := range items {
		if values, ok := item.([]int); ok {
			for _, value := range values {
				s.m[value] = Empty
			}
		} else {
			s.m[item] = Empty
		}
	}
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) Delete(item interface{}) {
	delete(s.m, item)
}

func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}

func (s *Set) Range() {
	for value, _ := range s.m {
		println(value)
	}
}

func (s *Set) ToArray() (arr []interface{}) {
	for value, _ := range s.m {
		arr = append(arr, value)
	}
	return
}
