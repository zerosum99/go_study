package main

import (
	"fmt"
	inter "go_interface_test2/interfaceTest"
	st "go_interface_test2/structTest"
)

func main() {
	fmt.Println(" go interface test 2 ")
	u := st.User{"dahlmoon", 55}

	get(&u)

	s := st.Student{"nullmoon", 30}
	get(&s)

	var a interface{} = 1

	i := a       // a와 i 는 dynamic type, 값은 1
	j := a.(int) // j는 int 타입, 값은 1

	println(i) // 포인터주소 출력
	println(j) // 1 출력

}

func get(obj inter.Getter) {
	fmt.Println("### interface call ### ")
	fmt.Println("name = ", obj.GetName())
	fmt.Println("age = ", obj.GetAge())
}
