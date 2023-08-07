package goceng

import (
	"context"
	"github.com/ikhsan892/goceng/application/infrastructures"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

type App interface {
	DB() *pgx.Conn
	WithDBConnection(name string) *pgx.Conn
	Settings() Config
	Shutdown() error
}

func New() *Goceng {
	g := &Goceng{}
	g.Start()

	return g
}

type Goceng struct {
	pgConns      map[string]*pgx.Conn
	ctx          context.Context
	cfg          Config
	notification chan struct{}
}

func (g *Goceng) DB() *pgx.Conn {
	return g.WithDBConnection("default")
}

func (g *Goceng) WithDBConnection(name string) *pgx.Conn {
	return g.pgConns[name]
}

func (g *Goceng) initPostgreSQL(cfg Config) error {
	conns := make(map[string]*pgx.Conn)

	var err error
	conns["default"], err = infrastructures.NewPostgreSQL(g.ctx, cfg.DB.URL)

	g.pgConns = conns

	return err
}

func (g *Goceng) Settings() Config {
	return g.cfg
}

// starting point of application
func (g *Goceng) Start() {
	ctx := context.Background()
	cfg := newConfig()

	g.ctx = ctx
	g.notification = make(chan struct{}, 1)
	g.cfg = cfg

	if err := g.initPostgreSQL(cfg); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Println("application started")
}

// for shutdown
func (g *Goceng) Shutdown() error {
	close(g.notification)

	return g.pgConns["default"].Close(g.ctx)
}
