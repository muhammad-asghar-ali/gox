package main

import (
	"embed"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/config"
)

var (
	//go:embed migrate/*.sql
	src embed.FS
)

func main() {
	dir, err := iofs.New(src, "migrate")
	if err != nil {
		log.Fatal("error creating IOFS migration source:", err)
		return
	}

	migrations, err := migrate.NewWithSourceInstance(
		"iofs",
		dir,
		config.GetConfig().ConnectionURI(),
	)
	if err != nil {
		log.Fatal("unable to create migration instance:", err)
		return
	}

	if err := migrations.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("error applying migrations:", err)
		return
	}

	log.Println("Migrations applied successfully.")
}
