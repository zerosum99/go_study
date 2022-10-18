package generic

import (
	"golang.org/x/exp/constraints"
)

// 제너릭 타입을 제한타입을로 지정
// 정수와 실수를 지정
func Add[T constraints.Integer | constraints.Float](a, b T) T { // 두 수의 계산
	return a + b
}

// int와 int32는 둘 다 다 정의해야 함
type Integer interface {
	int8 | int16 | int32 | int | int64
}

func Mul[T Integer](a, b T) T { // 두 수의 계산
	return a * b
}

// int와 int32는 둘 다 다 정의해야 함
type Number interface {
	int8 | int16 | int32 | int | int64 | float32 | float64
}

func Product[T Number](a, b T) T { // 두 수의 계산
	return a * b
}
