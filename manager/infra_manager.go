package manager

import (
	"context"
	"fmt"
	"simple-bank/config"
	db "simple-bank/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type InfraManager interface {
	openConn() error
	NewQuery() *db.Queries
	NewCtx() context.Context
	NewPool() *pgxpool.Pool
}

type infraManager struct {
	query *db.Queries
	cfg   *config.Config
	ctx   context.Context
	pool  *pgxpool.Pool
}

func (i *infraManager) openConn() error {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", i.cfg.Username, i.cfg.Password, i.cfg.DbName, i.cfg.Host, i.cfg.Port)

	i.ctx = context.Background()

	pool, err := pgxpool.New(i.ctx, connString)
	if err != nil {
		panic(err)
	}

	i.query = db.New(pool)
	i.pool = pool
	return nil
}

func (i *infraManager) NewQuery() *db.Queries {
	return i.query
}

func (i *infraManager) NewCtx() context.Context {
	return i.ctx
}

func (i *infraManager) NewPool() *pgxpool.Pool {
	return i.pool
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{
		cfg: cfg,
	}
	if err := conn.openConn(); err != nil {
		return nil, err
	}

	return conn, nil
}
