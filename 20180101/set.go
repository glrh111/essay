package main

import (
	"strconv"
	"hash/fnv"
	"fmt"
)


/*
    基于hash table 的集合 set
 */

type Key struct {
	value interface{}
}

func NewKey(value interface{}) *Key {
	return &Key{value}
}

func (k *Key) Value() interface{} {
	return k.value
}

func (k *Key) IsEqual(anotherKey *Key) bool {
	return k.Value() == anotherKey.Value()
}

type Set struct {
	keyList []*Key
	size int                // 实际存储的键值対的数量
	m int                   // 数组总大小
	minM int                // 设置的最小的数组的大小
}

func NewSet(initM int) *Set {
	minM := 10
	if initM < minM {
		initM = minM
	}
	return &Set{
		make([]*Key, initM),
		0,
		initM,
		minM,
	}
}

func NewSetFromList(lst []interface{}) (s *Set) {
	s = NewSet(10)
	for _, value := range lst {
		s.Add(value)
	}
	return
}

// resize array
func (b *Set) resize(cap int) {
	that := NewSet(cap)
	for i := 0; i < b.m; i++ {
		thisKey := b.keyList[i]
		if thisKey != nil {
			that.Add(thisKey.Value())
		}
	}
	*b = *that
}

// hash value of key
func (b *Set) hashIndex(item interface{}) (index int) {

	repString := ""
	switch item.(type) {
	case string:
		repString, _ = item.(string)
	case int:
		tempString, _ := item.(int)
		repString = strconv.Itoa(tempString)
	default:
		panic("type not support")
	}

	h := fnv.New32()
	h.Write([]byte(repString))
	index = int(h.Sum32() % uint32(b.m))
	return
}

func (b *Set) nextIndex(nowIndex int) int {
	if nowIndex >= b.m - 1 {
		return 0
	} else {
		return nowIndex + 1
	}
}

func (b *Set) Add(item interface{}) (ifAdd bool) {
	// 插入之前，判断 size
	if b.size > b.m / 2 {
		b.resize(2 * b.m)
	}

	hashIndex := b.hashIndex(item)
	key := NewKey(item)

	for {
		thisKey := b.keyList[hashIndex]

		if thisKey == nil { // 没值，设置新值
			b.keyList[hashIndex] = key
			b.size++
			ifAdd = true
			break
		} else {

			if key.IsEqual(thisKey) { // 命中, 什么也不干
				ifAdd = false
				break
			} else {                          // hashIndex + 1
				hashIndex = b.nextIndex(hashIndex)
			}
		}
	}
	return
}

// 是否包含某个元素
func (b *Set) Contains(item interface{}) (contains bool) {

	key := NewKey(item)

	hashIndex := b.hashIndex(item)

	for {
		thisKey := b.keyList[hashIndex]

		if thisKey == nil { // 没值，说明没找到
			contains = false
			break
		} else {

			if key.IsEqual(thisKey) { // 命中
				contains = true
				break
			} else {                   // hashIndex + 1
				hashIndex = b.nextIndex(hashIndex)
			}
		}
	}
	return
}

func (b *Set) IsEmpty() bool {
	return 0 == b.Size()
}

func (b *Set) Size() int {
	return b.size
}

func (b *Set) IteratorChan() chan interface{} {
	nowSize := 0
	i := 0
	c := make(chan interface{})
	go func () {
		for {
			if !(i<b.m && nowSize<b.size) {
				break
			}

			thisKey := b.keyList[i]

			if thisKey == nil {
				i++
				continue
			}
			nowSize++
			i++
			c <- thisKey.Value()
		}
		close(c)
	}()
	return c
}

func (s *Set) ToString() (str string) {
	str = fmt.Sprintf("Set: %d\n%s\n", s.Size(), s.ToLineString())
	return
}

func (s *Set) ToLineString() (str string) {
	str = ""
	for value := range s.IteratorChan() {
		tmpStr := ""
		switch value.(type) {
		case string:
			tmpStr, _ = value.(string)
		case int:
			tmpInt, _ := value.(int)
			tmpStr = strconv.Itoa(tmpInt)
		default:
			panic("type not support!")
		}
		str += fmt.Sprintf("%s ", tmpStr)
	}
	return
}

func (s *Set) ToList() (lst []interface{}) {
	lst = make([]interface{}, s.Size())
	index := 0
	for value := range s.IteratorChan() {
		lst[index] = value
		index++
	}
	return
}

/*
    下面是几个集合的操作
    UnionSet() ∪
    IntersectSet() ∩
 */
func UnionSet(setList ...*Set) (s *Set) {
	s = NewSet(10)
	for _, set := range setList {
		c := set.IteratorChan()
		for value := range c {
			s.Add(value)
		}
	}
	return
}

func IntersectSet(setList ...*Set) (s *Set) {
	s = NewSet(10)
	minSetIndex := 0
	minSetSize := setList[0].Size()
	if len(setList) > 1 {
		for setIndex, set := range setList[1:] {
			thisSetSize := set.Size()
			if minSetSize > thisSetSize {
				minSetSize = thisSetSize
				minSetIndex = setIndex
			}
		}
	}

	for value := range setList[minSetIndex].IteratorChan() {
		flag := true
		for _, set := range setList {
			if !set.Contains(value) {
				flag = false
				break
			}
		}
		if flag {
			s.Add(value)
		}
	}

	return
}

// s1 - s2
func SubtractSet(s1 *Set, s2 *Set) (s *Set) {
	s = NewSet(10)
	for value := range s1.IteratorChan() {
		if !s2.Contains(value) {
			s.Add(value)
		}
	}
	return
}
