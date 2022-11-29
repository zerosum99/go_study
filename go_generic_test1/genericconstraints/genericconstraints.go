package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	fmt.Println(" type constraints process ")

	var s1 Mystring = "hello"
	var s2 Mystring = "world"
	fmt.Println(Equal(s1, s2))
}

type ComparableHasher interface { // 타입 제한자 정의하기
	comparable    // == !=
	Hash() uint32 // hash 값 처리
}

type Mystring string

func (m Mystring) Hash() uint32 {
	h := fnv.New32a()
	h.Write([]byte(m))
	return h.Sum32()
}

func Equal[T ComparableHasher](a, b T) bool {
	if a == b {
		return true
	}
	return a.Hash() == b.Hash()
}
