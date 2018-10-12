package queue

type StrQueue struct {
	contents []string
}

func NewStrQueue(contents []string) *StrQueue {
	return &StrQueue{contents}
}

func (sq *StrQueue) Add(s string) {
	sq.contents = append(sq.contents, s)
}

func (sq *StrQueue) Empty() bool {
	return len(sq.contents) == 0
}

func (sq *StrQueue) Pop() string {
	if sq.Empty() {
		panic("Cannot pop from empty queue")
	}

	s := sq.contents[0]
	sq.contents = sq.contents[1:]
	return s
}

func (sq *StrQueue) Push(s string) {
	sq.contents = append(sq.contents, s)
}