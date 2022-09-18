package interface_test

type Get interface {
	GetName() string
}

type People struct {
	Name string
}

func (p People) GetName() string {
	return p.Name
}

type Animal struct {
	Name string
}

func (a Animal) GetName() string {
	return a.Name
}
