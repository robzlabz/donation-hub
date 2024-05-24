package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	Config         ApiConfig
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
		Config:         config,
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

func (a *API) Start() error {
	err := a.NetHttp.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
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
	page, limit := GetPageLimit(r)

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
	page, limit := GetPageLimit(r)

	status := r.URL.Query().Get("status")

	projects, err := a.ProjectService.ListProject(r.Context(), limit, page, status)
	if err != nil {
		ErrBadRequest(w, err.Error())
		return
	}

	SuccessResponse(w, projects)
}

func (a *API) HandlePostProjects(w http.ResponseWriter, r *http.Request) {
	var req reqRegister
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		ErrBadRequest(w, err.Error())
		return
	}

	//
	//// todo: handle disini
	//userID := int64(1) // todo: get from jwt
	//
	//input := request.ProjectRequestBody{
	//	Title:       req.Title,
	//	Description: req.Description,
	//	ImageURLs:   req.ImageUrls,
	//	DueAt:       req.DueAt,
	//	Currency:    req.Currency,
	//}
	//
	//err = a.ProjectService.SubmitProject(r.Context(), input)
	//if err != nil {
	//	ErrBadRequest(w, err.Error())
	//	return
	//}
	//
	//SuccessResponse(w, project)
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
