package storage

import (
	"context"
	"log"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresDB struct {
	sync.Mutex
	Pool *pgxpool.Pool
	Ctx  context.Context
	conn string
}

func NewPostgresDB(ctx context.Context, connectionString string) *PostgresDB {
	return &PostgresDB{
		Ctx:  ctx,
		conn: connectionString,
	}
}

func (db *PostgresDB) Connect() bool {
	db.Lock()
	if db.Pool == nil {
		log.Printf("Connect to db...")
		pool, err := pgxpool.Connect(db.Ctx, db.conn)
		if err != nil {
			db.Unlock()
			log.Printf("Unable to connect to database: %v\n", err)
			return false
		}
		db.Pool = pool
	}
	db.Unlock()
	return true
}

//func (db *PostgresDB) Connected() bool {
//	if db.Pool == nil {
//		if !connect() {
//			return false
//		}
//	}
//	return true
//}
