package stack

type StrStack struct {
	contents []string
}

func NewStrStack(contents []string) *StrStack {
	return &StrStack{contents}
}

func (ss *StrStack) Empty() bool {
	return len(ss.contents) == 0
}

func (ss *StrStack) Push(s string) {
	ss.contents = append(ss.contents, s)
}

func (ss *StrStack) Pop() string {
	if ss.Empty() {
		panic("Can't pop from empty stack")
	}

	s := ss.contents[len(ss.contents)-1]
	ss.contents = ss.contents[0:len(ss.contents)-1]
	return s
}
