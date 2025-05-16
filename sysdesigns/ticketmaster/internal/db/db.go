package db

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/config"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
)

var (
	instance *pgx.Conn
	once     sync.Once

	queries *entities.Queries
	qoc     sync.Once
)

func Connect() *pgx.Conn {
	once.Do(func() {
		conn, err := pgx.Connect(context.Background(), config.GetConfig().ConnectionURI())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		instance = conn
	})

	return instance
}

func Get() *pgx.Conn {
	if instance == nil {
		return Connect()
	}

	return instance
}

func Ping(ctx context.Context) error {
	return instance.Ping(ctx)
}

func Queries() *entities.Queries {
	qoc.Do(func() {
		conn := Get()
		if conn == nil {
			conn = Connect()
		}

		queries = entities.New(conn)
	})

	return queries
}
