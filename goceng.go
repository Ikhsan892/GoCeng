package goceng

import (
	"context"
	"github.com/ikhsan892/goceng/application/infrastructures"
	"github.com/jackc/pgx/v5"
	"github.com/meilisearch/meilisearch-go"
	"go.uber.org/zap"
	"log"
	"os"
)

type App interface {
	DB() *pgx.Conn
	SearchEngine() *meilisearch.Client
	WithDBConnection(name string) *pgx.Conn
	Settings() Config
	Shutdown() error
	ZapLogger() *zap.Logger
}

func New() *Goceng {
	g := &Goceng{}
	g.Start()

	return g
}

type Goceng struct {
	pgConns      map[string]*pgx.Conn
	searchEngine *meilisearch.Client
	ctx          context.Context
	cfg          Config
	logger       *zap.Logger
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

func (g *Goceng) initMeilisearch(cfg Config) {

	g.searchEngine = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:    "http://127.0.0.1:7700",
		APIKey:  "",
		Timeout: 0,
	})

}

func (g *Goceng) Settings() Config {
	return g.cfg
}

func (g Goceng) ZapLogger() *zap.Logger {
	return g.logger
}

func (g Goceng) SearchEngine() *meilisearch.Client {
	return g.searchEngine
}

// starting point of application
func (g *Goceng) Start() {
	ctx := context.Background()
	cfg := newConfig()

	g.ctx = ctx
	g.notification = make(chan struct{}, 1)
	g.cfg = cfg
	g.logger = NewZapLogger()

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
