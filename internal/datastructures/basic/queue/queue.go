package queue

type Queue struct {
	Data   []interface{}
	length int
}

func New() *Queue {
	return &Queue{
		Data:   nil,
		length: 0,
	}
}

func (q *Queue) Enqueue(v interface{}) {
	q.Data = append(q.Data, v)
	q.length++
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	value := q.Data[0]

	copy(q.Data, q.Data[1:])  // Shift elements to the left
	q.Data[len(q.Data)-1] = 0 // Zero out the last element
	q.Data = q.Data[:len(q.Data)-1]
	q.length--

	return value
}

func (q *Queue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.Data[0]
}

func (q *Queue) Size() int {
	return q.length
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
