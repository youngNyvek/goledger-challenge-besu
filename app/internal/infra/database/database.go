package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func MustConnectPostgres() *sql.DB {
	host := envOr("POSTGRES_HOST", "localhost")
	port := envOr("POSTGRES_PORT", "5432")
	user := envOr("POSTGRES_USER", "admin")
	pass := envOr("POSTGRES_PASSWORD", "admin123")
	dbname := envOr("POSTGRES_DB", "goledger")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil { log.Fatalf("Erro abrindo PostgreSQL: %v", err) }
	if err := db.Ping(); err != nil { log.Fatalf("Erro conectando PostgreSQL: %v", err) }
	return db
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" { return v }
	return def
}
