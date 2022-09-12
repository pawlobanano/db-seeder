package seed

import (
	"context"
	"embed"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var DataFS embed.FS

func RunSeeds() {
	files, err := DataFS.ReadDir("data")
	if err != nil {
		log.Fatal("error reading seed directory", err)
	}

	for _, file := range files {
		seeder(context.Background(), dbConn(), file.Name())
	}
}

func seeder(ctx context.Context, pool *pgxpool.Pool, seedName string) {
	query, err := DataFS.ReadFile(fmt.Sprintf("data/%s", seedName))
	if err != nil {
		panic(err)
	}

	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		log.Println("a non-default isolation level is used that the driver doesn't support")
	}

	_, err = tx.Exec(ctx, string(query))
	if err != nil {
		log.Printf("doing rollback of %s", seedName)
		tx.Rollback(ctx)
		return
	}

	tx.Commit(ctx)
	log.Printf("%s loaded into the database", seedName)
}

func dbConn() *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), "postgres://postgres:password@localhost:5432/db_instance_name?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal("error pinging db connection", err)
	}

	return pool
}
