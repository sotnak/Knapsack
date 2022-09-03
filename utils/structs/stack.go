package structs

type Stack[T any] struct {
	head   *Node[T]
	length int
}

type Node[T any] struct {
	elem func()
	next *Node[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil, 0}
}

func (this *Stack[T]) Push(elem func()) {
	node := &Node[T]{elem, this.head}

	this.head = node
	this.length++
}

func (this *Stack[T]) Pop() func() {
	if this.length == 0 {
		panic("empty stack")
	}

	newHead := this.head.next
	ret := this.head.elem

	this.head = newHead
	this.length--

	return ret
}

func (this *Stack[T]) Empty() bool {
	return this.length <= 0
}
