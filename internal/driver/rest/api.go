package rest

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/isdzulqor/donation-hub/internal/core/entity"
	"github.com/isdzulqor/donation-hub/internal/core/service/project"
	"github.com/isdzulqor/donation-hub/internal/core/service/user"
	"github.com/isdzulqor/donation-hub/internal/driver/request"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
)

type API struct {
	DB             *sqlx.DB
	UserService    user.Service
	ProjectService project.Service
}

var httpSuccess = HttpSuccess{}
var httpError = HttpError{}

func (a *API) HandlePing(w http.ResponseWriter, r *http.Request) {
	type PingPong struct {
		Ping string `json:"ping"`
	}
	pong := PingPong{Ping: "pong"}
	httpSuccess.SuccessResponse(w, pong)
}

func (a *API) LogRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: %s %s", r.Method, r.URL.Path)
}

func (a *API) HandlePostRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req request.RegisterRequestBody
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			httpError.ErrBadRequest(w, err.Error())
			return
		}

		registerUser, err := a.UserService.RegisterUser(r.Context(), req)
		if err != nil {
			httpError.ErrBadRequest(w, err.Error())
			return
		}

		responseRegister := struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
		}{
			ID:       registerUser.ID,
			Username: registerUser.Username,
			Email:    registerUser.Email,
		}

		httpSuccess.SuccessResponse(w, responseRegister)
	}
}

func (a *API) HandlePostLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse the request body
		var req request.LoginRequestBody
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			httpError.ErrBadRequest(w, err.Error())
			return
		}

		// Call the LoginUser method from the userService
		loginUser, accessToken, err := a.UserService.LoginUser(r.Context(), req)
		if err != nil {
			httpError.ErrUnauthorized(w, err.Error())
			return
		}

		loginResponse := struct {
			ID          int64  `json:"id"`
			Email       string `json:"email"`
			Username    string `json:"username"`
			AccessToken string `json:"access_token"`
			Ts          int64  `json:"ts"`
		}{
			ID:          loginUser.ID,
			Email:       loginUser.Email,
			Username:    loginUser.Username,
			AccessToken: accessToken,
		}

		httpSuccess.SuccessResponse(w, loginResponse)
	} else {
		httpError.ErrNotFound(w)
	}
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
	httpSuccess := HttpSuccess{}
	httpSuccess.SuccessResponse(w, users)
}

func (a *API) HandleGetProjects(w http.ResponseWriter, r *http.Request) {
	httpSuccess.SuccessResponse(w, struct {
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
