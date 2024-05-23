package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/isdzulqor/donation-hub/internal/core/entity"
	"github.com/isdzulqor/donation-hub/internal/core/service/project"
	"github.com/isdzulqor/donation-hub/internal/core/service/user"
	"github.com/isdzulqor/donation-hub/internal/driver/request"
	"github.com/isdzulqor/donation-hub/internal/utils/validator"
	"github.com/jmoiron/sqlx"
)

type ApiConfig struct {
	DB             *sqlx.DB
	UserService    user.Service
	ProjectService project.Service
}

type API struct {
	config         ApiConfig
	NetHttp        *http.Server
	UserService    user.Service
	ProjectService project.Service
}

func NewAPI(config ApiConfig) (*API, error) {
	err := validator.Validate().Struct(config)
	if err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	app := &API{
		config:         config,
		UserService:    config.UserService,
		ProjectService: config.ProjectService,
	}

	// make http handler with net/http
	http.HandleFunc("GET /ping", app.HandlePing)
	http.HandleFunc("POST /register", app.HandlePostRegister)
	http.HandleFunc("POST /login", app.HandlePostLogin)
	http.HandleFunc("GET /users", app.HandleGetUsers)
	http.HandleFunc("GET /projects", app.HandleGetProjects)
	http.HandleFunc("POST /projects", app.HandlePostProjects)
	http.HandleFunc("POST /projects/{id}/review", app.HandleProjectReview)
	http.HandleFunc("GET /projects/{id}", app.HandleProjectDetails)
	http.HandleFunc("GET /projects/{id}/donate", app.HandlePostProjectDonation)

	app.NetHttp = &http.Server{}

	return app, nil
}

func (a *API) HandlePing(w http.ResponseWriter, r *http.Request) {
	type PingPong struct {
		Ping string `json:"ping"`
	}
	pong := PingPong{Ping: "pong"}
	SuccessResponse(w, pong)
}

func (a *API) LogRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: %s %s", r.Method, r.URL.Path)
}

func (a *API) HandlePostRegister(w http.ResponseWriter, r *http.Request) {
	var req user.InputRegister
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		ErrBadRequest(w, err.Error())
		return
	}

	registerUser, err := a.UserService.RegisterUser(r.Context(), req)
	if err != nil {
		ErrBadRequest(w, err.Error())
		return
	}

	SuccessResponse(w, respRegister{
		ID:       registerUser.ID,
		Username: registerUser.Username,
		Email:    registerUser.Email,
	})
}

func (a *API) HandlePostLogin(w http.ResponseWriter, r *http.Request) {
	var req request.LoginRequestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		ErrBadRequest(w, err.Error())
		return
	}

	// Call the LoginUser method from the userService
	loginUser, accessToken, err := a.UserService.LoginUser(r.Context(), req)
	if err != nil {
		ErrUnauthorized(w, err.Error())
		return
	}

	SuccessResponse(w, respSuccessLogin{
		ID:          loginUser.ID,
		Email:       loginUser.Email,
		Username:    loginUser.Username,
		AccessToken: accessToken,
	})
}

func (a *API) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	// Get the page and limit values from the query parameters
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	// Default the page and limit to 1 and 10 respectively if they are not provided
	var page, limit int
	if pageStr == "" {
		page = 1
	} else {
		page, _ = strconv.Atoi(pageStr) // error handling omitted for brevity
	}
	if limitStr == "" {
		limit = 10
	} else {
		limit, _ = strconv.Atoi(limitStr) // error handling omitted for brevity
	}

	users, err := a.UserService.GetListUser(r.Context(), limit, page, entity.UserRoleDonor)
	fmt.Println(users)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the response struct to JSON
	SuccessResponse(w, users)
}

func (a *API) HandleGetProjects(w http.ResponseWriter, r *http.Request) {
	SuccessResponse(w, struct {
		Message string `json:"message"`
	}{
		Message: "If you see this, you're authorized to access this route.",
	})
}

func (a *API) HandlePostProjects(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Projects"))
}

func (a *API) HandleProjectReview(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Project Review"))
}

func (a *API) HandleProjectDetails(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Project Details"))
}

func (a *API) HandlePostProjectDonation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Project Donation"))
}

func (a *API) HandleGetProjectDonation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Project Donation"))
}
