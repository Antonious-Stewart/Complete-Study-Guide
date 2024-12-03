package linked_list

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type LinkedList[T any] struct {
	Head   *Node[T]
	Tail   *Node[T]
	length int
}

func (l *LinkedList[T]) New() *LinkedList[T] {
	return &LinkedList[T]{
		Head: nil,
		Tail: nil,
	}
}

func (l *LinkedList[T]) GetHead() *Node[T] {
	if l.length == 0 {
		return nil
	}

	return l.Head
}
