package structs

type Stack struct {
	head   *Node
	length int
}

type Node struct {
	elem func()
	next *Node
}

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (this *Stack) Push(elem func()) {
	node := &Node{elem, this.head}

	this.head = node
	this.length++
}

func (this *Stack) Pop() func() {
	if this.length == 0 {
		panic("empty stack")
	}

	newHead := this.head.next
	ret := this.head.elem

	this.head = newHead
	this.length--

	return ret
}

func (this *Stack) Empty() bool {
	return this.length <= 0
}
