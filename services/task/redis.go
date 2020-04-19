package task

type redis struct {

}

func (rd *redis) pull()  {

}

func (rd *redis) push()  {

}

func (rd *redis) run() queue {
	return new(redis)
}