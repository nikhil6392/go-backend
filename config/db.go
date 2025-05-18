package config

import (
	"context"
	"os"
	"log"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)



var DB *pgxpool.Pool


func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}
	fmt.Println("Connecting to DB with DSN:")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v\n", err) // FIXED
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to ping DB: %v\n", err) // FIXED
	}

	DB = pool
	fmt.Println("✅ Connected to Neon PostgreSQL")
}
