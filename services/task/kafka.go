package task

import "chat/tool"

type kafka struct {

}

func (kfk *kafka) pull()  {

}

func (kfk *kafka) push()  {

}

func (kfk *kafka) run() {
	tool.Logger.Info().Msg("this is kafka")
}