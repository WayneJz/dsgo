package set

import (
	"fmt"
	"strings"
)

// Set 集合
type Set map[interface{}]struct{}

// NewSet 新建一个集合
func NewSet() *Set {
	s := Set(make(map[interface{}]struct{}))
	return &s
}

// CreateSet 新建一个集合并赋初始元素
func CreateSet(elements ...interface{}) *Set {
	s := NewSet()
	s.Add(elements...)
	return s
}

// Add 向集合内添加元素
func (s *Set) Add(elements ...interface{}) {
	for _, element := range elements {
		(*s)[element] = struct{}{}
	}
}

// Remove 删除集合内元素
func (s *Set) Remove(elements ...interface{}) {
	for _, element := range elements {
		delete(*s, element)
	}
}

// Clear 清空集合
func (s *Set) Clear() {
	for k := range *s {
		delete(*s, k)
	}
}

// Contains 判断集合是否包含此元素
func (s *Set) Contains(element interface{}) bool {
	_, ok := (*s)[element]
	return ok
}

// Members 获取集合内所有元素
func (s *Set) Members() []interface{} {
	res := []interface{}{}
	for k := range *s {
		res = append(res, k)
	}
	return res
}

// DeepCopy 深拷贝集合
func (s *Set) DeepCopy() *Set {
	res := NewSet()
	for k := range *s {
		(*res)[k] = struct{}{}
	}
	return res
}

// Union 求集合s与集合set的并集
func (s *Set) Union(set *Set) *Set {
	res := s.DeepCopy()
	for k := range *set {
		(*res)[k] = struct{}{}
	}
	return res
}

// Intersection 求集合s与集合set的交集
func (s *Set) Intersection(set *Set) *Set {
	res := s.DeepCopy()
	removes := []interface{}{}
	for k := range *set {
		if !res.Contains(k) {
			removes = append(removes, k)
		}
	}
	res.Remove(removes...)
	return res
}

// Difference 求集合s与集合set的差集 ({x∣x∈s,且x∉set})
func (s *Set) Difference(set *Set) *Set {
	res := s.DeepCopy()
	removes := []interface{}{}
	for k := range *set {
		if res.Contains(k) {
			removes = append(removes, k)
		}
	}
	res.Remove(removes...)
	return res
}

// IsSubset 判断集合s是否为集合set的子集
func (s *Set) IsSubset(set *Set) bool {
	for k := range *s {
		if _, ok := (*set)[k]; !ok {
			return false
		}
	}
	return true
}

// IsEqual 判断集合s是否与集合set相等
func (s *Set) IsEqual(set *Set) bool {
	return s.IsSubset(set) && set.IsSubset(s)
}

// String 打印方法实现
func (s *Set) String() string {

	builder := &strings.Builder{}
	for k := range *s {
		builder.WriteString(fmt.Sprintf("%v ", k))
	}
	return fmt.Sprintf("Set{ %s}", builder.String())
}
