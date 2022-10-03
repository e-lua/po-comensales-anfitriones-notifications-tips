package models

import (
	"context"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var PostgresCN = Conectar_Pg_DB()

var (
	once_pg sync.Once
	p_pg    *pgxpool.Pool
)

func Conectar_Pg_DB() *pgxpool.Pool {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer cancelara el contexto
	defer cancel()

	once_pg.Do(func() {
		//urlString := "postgres://postgresxv6:4kdfghklfg1463hadgkj45345M@postgres-master:5432/postgresxv6?pool_max_conns=150"
		urlString := "postgres://postgresxv7:4kdfghklfg1463hadgkj45345M@postgres-master:5432/postgresxv7?pool_max_conns=10"
		config, _ := pgxpool.ParseConfig(urlString)
		p_pg, _ = pgxpool.ConnectConfig(ctx, config)
	})

	return p_pg
}
