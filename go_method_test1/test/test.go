package test

func Add(x, y int) int {
	return x + y
}

type Arth struct {
	X int
	Y int
}

func (ar Arth) Add() int {
	return ar.X + ar.Y
}

type Adder interface {
	Add() int
}
