package main

import (
	"fmt"
	"github.com/isdzulqor/donation-hub/internal/core/service/project"
	"github.com/isdzulqor/donation-hub/internal/core/service/user"
	"github.com/isdzulqor/donation-hub/internal/driven/storage/mysql/projectstr"
	"github.com/isdzulqor/donation-hub/internal/driven/storage/mysql/userstr"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"

	"github.com/isdzulqor/donation-hub/internal/driver/rest"
)

func main() {
	connectionString := "root:@tcp(localhost:3306)/donation_hub"
	db, err := ConnectToDatabase(connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	storageUser := userstr.New(userstr.Config{SQLClient: db})
	userService := user.NewService(storageUser)

	projectStorage := projectstr.New(projectstr.Config{SQLClient: db})
	projectService := project.NewService(projectStorage)

	var restApi = rest.API{
		DB:             db,
		UserService:    userService,
		ProjectService: projectService,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", restApi.HandlePing)

	mux.HandleFunc("/users/register", restApi.HandlePostRegister)
	mux.HandleFunc("POST /users/login", restApi.HandlePostLogin)
	mux.HandleFunc("/users", restApi.HandleGetUsers)
	mux.HandleFunc("/projects", restApi.HandleGetProjects)
	mux.HandleFunc("POST /projects", restApi.HandlePostProjects)
	mux.HandleFunc("PUT /projects/{id}/review", restApi.HandleProjectReview)
	mux.HandleFunc("/projects/{id}", restApi.HandleProjectDetails)
	mux.HandleFunc("POST /projects/{id}/donation", restApi.HandlePostProjectDonation)
	mux.HandleFunc("/projects/{id}/donation", restApi.HandleGetProjectDonation)

	log.Fatal(http.ListenAndServe(":8180", mux))
}

func ConnectToDatabase(connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
