package rest

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/isdzulqor/donation-hub/internal/core/entity"
	"github.com/isdzulqor/donation-hub/internal/core/service/project"
	"github.com/isdzulqor/donation-hub/internal/core/service/user"
	"github.com/isdzulqor/donation-hub/internal/driver/request"
	"github.com/isdzulqor/donation-hub/internal/driver/response"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
	"time"
)

type API struct {
	DB             *sqlx.DB
	UserService    user.Service
	ProjectService project.Service
}

func (a *API) HandlePing(w http.ResponseWriter, r *http.Request) {
	// write json pong
	w.Header().Set("Content-Type", "application/json")
	type PingPong struct {
		Ping string `json:"ping"`
	}
	pong := PingPong{Ping: "pong"}
	res, err := json.Marshal(pong)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (a *API) LogRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: %s %s", r.Method, r.URL.Path)
}

func (a *API) HandlePostRegister(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterRequestBody
	log.Printf("RequesTTT: %s %s", r.Method, r.URL.Path)
	err := json.NewDecoder(r.Body).Decode(&req)
	log.Printf("Request: %s", req)
	if err != nil {
		log.Println(err)
		// todo : response error properly
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// print request body
	log.Println(req)

	err = a.UserService.RegisterUser(r.Context(), req)
	if err != nil {
		log.Println(err)
		// todo : response error properly
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	r.Header.Add("Content-Type", "application/json")
	log.Printf("Response: %s", `{"ok":true}`)
	w.Write([]byte(`{"ok":true}`))
}

func (a *API) HandlePostLogin(w http.ResponseWriter, r *http.Request) {

	// Parse the request body
	var req request.LoginRequestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the LoginUser method from the userService
	user, err := a.UserService.LoginUser(r.Context(), req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// If login is successful, return a JSON response with the user details
	res, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
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

	// Create a new ResponseListUser struct
	resp := response.ResponseBodySuccess{
		Ok:   true,
		Data: users,
		Ts:   time.Now().Unix(),
	}

	// Convert the response struct to JSON
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (a *API) HandleGetProjects(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Projects"))
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
