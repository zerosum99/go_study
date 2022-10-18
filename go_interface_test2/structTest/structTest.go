package structtest

type Student struct {
	Name string
	Age  int
}

func (s Student) GetName() string {
	return s.Name
}

func (s Student) GetAge() int {
	return s.Age
}

type User struct {
	Name string
	Age  int
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetAge() int {
	return u.Age
}
