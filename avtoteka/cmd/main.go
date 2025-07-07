package main

import (
	"avtoteka/avtoteka/config"
	"avtoteka/avtoteka/internal/handlers"
	"avtoteka/avtoteka/internal/repository"
	"avtoteka/avtoteka/internal/services"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func applyMigrations(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failes to create driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to init migrate: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	version, dirty, err := m.Version()
	if err != nil {
		return fmt.Errorf("failed to get migration version: %w", err)
	}
	log.Printf("Migrations applied. Current version: %d (dirty: %v)", version, dirty)

	return nil
}

func main() {
	cfg := config.Load()
	db, err := repository.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := applyMigrations(db); err != nil {
		log.Fatalf("Migrations failed: %v", err)
	}
	defer db.Close()

	carRepo := repository.NewCarRepository(db)
	carService := services.NewCarService(carRepo)
	carHandler := handlers.NewCarHandler(carService)

	r := mux.NewRouter()

	r.HandleFunc("/api/cars", carHandler.CreateCar).Methods("POST")
	r.HandleFunc("/api/cars/{id}", carHandler.GetCar).Methods("GET")
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) { //check work
		w.Write([]byte("pong"))
	})

	err = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return err
		}

		methods, err := route.GetMethods()
		if err != nil {
			return nil // Пропускаем если нет методов
		}

		log.Printf("Registered route: %-6s %s", methods, path)
		return nil
	})

	if err != nil {
		log.Printf("Route logging error: %v", err)
	}

	log.Printf("Server is run on http://localhost:%s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
}
