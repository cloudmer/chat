package task

type queue interface {
	pull()
	push()
}

