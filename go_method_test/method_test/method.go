package method_test

type User struct {
	Name string
	Age  string
}

func (u User) Get() string {
	return u.Name + u.Age
}
