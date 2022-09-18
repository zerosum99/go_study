package main

//다른 모듈을 가져오려면 .mod 에 require와 replace를 갱신해야 함

import (
	"fmt"

	inter "go_interface_test/interface_test"

	mt "aaa.com/go_method_test/method_test"
)

func main() {
	fmt.Println(" go interface ")
	p := inter.People{"ddd"}
	fmt.Println(" peopel get name ", GetName(p))
	a := inter.Animal{"aaa"}
	fmt.Println(" animal get name ", GetName(a))

	//다른 모듈을 import 처리하기

	u := mt.User{"가가가", "77"}
	fmt.Println("User ", u)

}

// 인터페이스 타입으로 지정한 함수...
// 실제 인터페이스 내의 메소드가 구현되면 자동으로 인터페이스 타입으로 연결됨
func GetName(x inter.Get) string {
	return x.GetName()
}
