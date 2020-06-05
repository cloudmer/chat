package task

import "chat/tool"

type redis struct {

}

func (rd *redis) pull()  {

}

func (rd *redis) push()  {

}

func (rd *redis) run() {
	tool.Logger.Info().Msg("this is redis")
}