// Package main is the entry point of the application.
package main

import (
	"fmt"
	"github.com/isdzulqor/donation-hub/internal/core/service/project"
	"github.com/isdzulqor/donation-hub/internal/core/service/user"
	"github.com/isdzulqor/donation-hub/internal/driven/storage/mysql/projectstr"
	"github.com/isdzulqor/donation-hub/internal/driven/storage/mysql/userstr"
	"github.com/isdzulqor/donation-hub/internal/driver/middleware/jwt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"

	"github.com/isdzulqor/donation-hub/internal/driver/rest"
	"github.com/joho/godotenv"
)

// main is the main function of the application.
// It sets up the database connection, initializes the services and starts the HTTP server.
func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Database connection string
	connectionString := getDatabaseConnectionString()

	// Connect to the database
	db, err := ConnectToDatabase(connectionString)
	if err != nil {
		panic(err)
	}

	// Get JWT secret key and issuer from environment variables
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	jwtIssuer := os.Getenv("JWT_ISSUER")

	// Initialize the user storage and service
	storageUser := userstr.New(userstr.Config{SQLClient: db})
	jwtService := encryption.NewJWTService(jwtSecretKey, jwtIssuer)
	userService := user.NewService(storageUser, jwtService)

	// Initialize the project storage and service
	projectStorage := projectstr.New(projectstr.Config{SQLClient: db})
	projectService := project.NewService(projectStorage)

	// Initialize the REST API
	var restApi = rest.API{
		DB:             db,
		UserService:    userService,
		ProjectService: projectService,
	}

	// Set up the HTTP server
	mux := http.NewServeMux()

	// Define the HTTP routes
	mux.HandleFunc("/ping", restApi.HandlePing)
	mux.HandleFunc("/users/register", restApi.HandlePostRegister)
	mux.HandleFunc("/users/login", restApi.HandlePostLogin)
	mux.HandleFunc("/projects/{id}", restApi.HandleProjectDetails)

	// can be public or private
	mux.Handle("/projects", jwtService.Middleware(http.HandlerFunc(restApi.HandleGetProjects), true))

	// Protected routes
	mux.Handle("/users", jwtService.Middleware(http.HandlerFunc(restApi.HandleGetUsers), false))
	mux.Handle("POST /projects", jwtService.Middleware(http.HandlerFunc(restApi.HandlePostProjects), false))
	mux.Handle("/projects/{id}/review", jwtService.Middleware(http.HandlerFunc(restApi.HandleProjectReview), false))
	mux.Handle("/projects/{id}/donation", jwtService.Middleware(http.HandlerFunc(restApi.HandlePostProjectDonation), false))

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8180", mux))
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

// getDatabaseConnectionString constructs a database connection string using environment variables.
// It retrieves the database username, password, host, port, and name from the environment variables
// DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, and DB_NAME respectively.
// It then formats these values into a connection string in the format "username:password@tcp(host:port)/dbname".
// The function returns the constructed connection string.
func getDatabaseConnectionString() string {
	// Retrieve the database username from the environment variable DB_USERNAME
	dbUsername := os.Getenv("DB_USERNAME")
	// Retrieve the database password from the environment variable DB_PASSWORD
	dbPassword := os.Getenv("DB_PASSWORD")
	// Retrieve the database host from the environment variable DB_HOST
	dbHost := os.Getenv("DB_HOST")
	// Retrieve the database port from the environment variable DB_PORT
	dbPort := os.Getenv("DB_PORT")
	// Retrieve the database name from the environment variable DB_NAME
	dbName := os.Getenv("DB_NAME")

	// Format the retrieved values into a connection string and return it
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
}
