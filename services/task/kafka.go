package task

type kafka struct {

}

func (kfk *kafka) pull()  {

}

func (kfk *kafka) push()  {

}

func (kfk *kafka) run() queue {
	return new(kafka)
}