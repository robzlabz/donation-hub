// Package main is the entry point of the application.
package main

import (
	"fmt"
	"log"

	"github.com/gosidekick/goconfig"
	"github.com/isdzulqor/donation-hub/internal/core/service/project"
	"github.com/isdzulqor/donation-hub/internal/core/service/user"
	"github.com/isdzulqor/donation-hub/internal/driven/storage/mysql/projectstr"
	"github.com/isdzulqor/donation-hub/internal/driven/storage/mysql/userstr"
	encryption "github.com/isdzulqor/donation-hub/internal/driver/middleware/jwt"
	"github.com/isdzulqor/donation-hub/internal/driver/rest"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	TokenIssuer                   string `cfg:"token_issuer" required:"true"`
	TokenSecret                   string `cfg:"token_secret" required:"true"`
	AccessTokenDurationInSeconds  int    `cfg:"access_token_duration_in_seconds" required:"true"`
	RefreshTokenDurationInSeconds int    `cfg:"refresh_token_duration_in_seconds" required:"true"`
	DatabaseUrl                   string `cfg:"database_url" required:"true" cfgDefault:"root@tcp(127.0.0.1)/donation_hub"`
	PhotoBucketName               string `cfg:"photo_bucket_name" required:"true"`
	Port                          int    `cfg:"port" required:"true"`
}

// main is the main function of the application.
// It sets up the database connection, initializes the services and starts the HTTP server.
func main() {
	cfg := Config{}
	err := goconfig.Parse(&cfg)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	fmt.Println(cfg)
	// Connect to the database
	db, err := ConnectToDatabase(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Initialize the user storage and service
	storageUser, err := userstr.New(userstr.Config{SQLClient: db})
	if err != nil {
		log.Fatalf("failed to initialize user storage: %v", err)
	}
	jwtService := encryption.NewJWTService(cfg.TokenSecret, cfg.TokenIssuer)
	userService := user.NewService(storageUser, jwtService)

	// Initialize the project storage and service
	projectStorage := projectstr.New(projectstr.Config{SQLClient: db})
	projectService := project.NewService(projectStorage)

	// Initialize the REST API
	_, err = rest.NewAPI(rest.ApiConfig{
		DB:             db,
		UserService:    userService,
		ProjectService: projectService,
	})

	if err != nil {
		log.Fatalf("failed to start app %v", err)
	}
}

// ConnectToDatabase connects to the database using the provided connection string.
// It returns a sqlx.DB object and an error.
func ConnectToDatabase(connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
