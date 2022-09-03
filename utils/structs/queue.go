package structs

type Queue[T any] struct {
	head   *queueNode[T]
	tail   *queueNode[T]
	length int
}

type queueNode[T any] struct {
	elem T
	prev *queueNode[T]
	next *queueNode[T]
}

func NewQueue[T any]() Container[T] {
	var container Container[T] = &Queue[T]{nil, nil, 0}

	return container
}

func (this *Queue[T]) Push(elem T) {

	var node *queueNode[T]

	if this.Empty() {
		node = &queueNode[T]{elem, nil, nil}
		this.tail = node
	} else {
		node = &queueNode[T]{elem, nil, this.head}
		this.head.prev = node
	}
	this.head = node
	this.length++
}

func (this *Queue[T]) Pop() T {
	if this.length == 0 {
		panic("empty stack")
	}

	newTail := this.tail.prev
	ret := this.tail.elem

	this.tail = newTail
	this.length--

	return ret
}

func (this *Queue[T]) Empty() bool {
	return this.length <= 0
}

func (this *Queue[T]) Len() int {
	return this.length
}
