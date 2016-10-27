package set

import (
	"fmt"
	"sort"
	"sync"
)

// set 对象
type Set struct {
	sync.RWMutex
	m map[int]bool
}

// 添加元素
func (s *Set) Add(items ...int) {
	s.Lock()
	defer s.Unlock()
	for _, v := range items {
		s.m[v] = true
	}
}

// 删除元素
func (s *Set) Remove(items ...int) {
	s.Lock()
	defer s.Unlock()
	for _, v = range items {
		delete(s.m, v)
	}
}

// 判断元素是否存在
func (s *Set) Has(items ...int) bool {
	s.RLock()
	defer s.RUnlock()
	for _, v := range items {
		if _, ok := s.m[v]; !ok {
			return false
		}
	}
	return true
}

// 统计元素个数
func (s *Set) Count() int {
	s.Lock()
	defer s.Unlock()
	return len(s.m)
}

// 清空集合
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
}

// 空集合判断
func (s *Set) Empty() bool {
	s.Lock()
	defer s.Unlock()
	return len(s.m) == 0
}

// 获取元素列表（无序）
func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func (s *Set) SortedList() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	sort.Ints(list)
	return list
}

// 新建集合对象
func New(items ...int) *Set {
	s := &Set{
		m: make(map[int]bool, len(items)),
	}
	s.Add(items...)
	return s
}
