package main

import (
	"github.com/ikhsan892/goceng"
	"github.com/ikhsan892/goceng/adapter"
	"github.com/ikhsan892/goceng/adapter/gui"
	"github.com/ikhsan892/goceng/adapter/meilisearch"
	"github.com/ikhsan892/goceng/adapter/web"
)

func main() {

	app := goceng.New()

	go adapter.RunAdapter(web.NewEcho(app))
	go adapter.RunAdapter(meilisearch.NewMeilisearchAdapter())

	// gui adapter must run in main goroutine
	adapter.RunAdapter(gui.NewFyneAdapter(app))

}
