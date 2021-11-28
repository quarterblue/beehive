package api

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.1"

type config struct {
	port   int
	env    string
	dbUser string
	dbPass string
}
type application struct {
	config config
	logger *log.Logger
}

func Server() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.dbUser, "user", "maxroach", "User name for DB")
	flag.StringVar(&cfg.dbPass, "pass", "", "User pass for DB")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/bee_hive?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	defer db.Close()
	fmt.Println("Database beehive_db opened and ready.")

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server.
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}
