package adapter

import "log"

type IAdapter interface {
	Init() error
}

func RunAdapter(adapter IAdapter) {
	if err := adapter.Init(); err != nil {
		log.Println(err)
	}
}
