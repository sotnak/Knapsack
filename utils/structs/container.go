package structs

type Container[T any] interface {
	Push(elem T)
	Pop() T
	Empty() bool
	Len() int
}
