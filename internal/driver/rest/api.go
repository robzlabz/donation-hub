package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	JWTSecret      string
}

type API struct {
	Config         ApiConfig
	Router         *http.ServeMux
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

	router := http.NewServeMux()

	// make http handler with net/http
	//router.HandleFunc("GET /", app.LogRequest)
	//router.HandleFunc("GET /ping", app.HandlePing)
	router.HandleFunc("/users/register", app.HandlePostRegister)
	router.HandleFunc("/users/login", app.HandlePostLogin)
	router.HandleFunc("/users", app.HandleGetUsers)
	router.HandleFunc("/projects", app.HandleGetProjects)
	//router.HandleFunc("/projects", app.JWTMiddleware(app.HandlePostProjects, true))
	router.HandleFunc("/projects/{id}/review", app.HandleProjectReview)
	router.HandleFunc("/projects/{id}", app.HandleProjectDetails)
	router.HandleFunc("/projects/{id}/donate", app.HandlePostProjectDonation)

	server := http.Server{
		Addr:    ":8180",
		Handler: router,
	}

	fmt.Printf("Go Runtime Version %s\n", runtime.Version())
	fmt.Println("Server is running on port 8180")
	err = server.ListenAndServe()
	if err != nil {
		return nil, err
	}

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
	// write it's working
	w.Write([]byte("It's working"))
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
	//handle jwt
	var req reqRegister
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		ErrBadRequest(w, err.Error())
		return
	}

	fmt.Println(req)

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

type CustomClaims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func (a *API) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(a.Config.JWTSecret), nil
	})
}

func (a *API) JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		// If the endpoint can be optional and there's no Authorization header, pass the request to the next handler
		//if canBeOptional && authHeader == "" {
		//	next.ServeHTTP(w, r)
		//	return
		//}

		// If the Authorization header does not contain "Bearer ", respond with an error
		if !strings.Contains(authHeader, "Bearer ") {
			ErrUnauthorized(w, "invalid access token")
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		token, err := a.ValidateToken(tokenString)

		// If there's an error in parsing the token or the token is not valid, respond with an error
		if err != nil || !token.Valid {
			ErrUnauthorized(w, "invalid access token")
			return
		}

		// If the token is valid, pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
