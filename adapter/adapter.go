package adapter

import "github.com/ikhsan892/goceng"

type Adapter interface {
	Init(app goceng.App) error
}
