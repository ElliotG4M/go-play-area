package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ElliotG4M/PlayArea/models"
)

// Defines properties/attributes of the 'class'
type userController struct {
	userIDPattern *regexp.Regexp
}

// Define a method for the 'class'
// uc userController binds the type userController to the method
// ServeHTTP is the method name with params for the request and response
// Because we've used the exact name ServeHTTP with these params, it automatically implements the Handler interface in the net/http package
// i.e. if we have an existing interface and we define a method in exactly the same way, Go will assume we are implementing said interface
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getUsers(w, r)
		case http.MethodPost:
			uc.addUser(w, r)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		switch r.Method {
		case http.MethodGet:
			uc.getUser(id, w)
		case http.MethodPut:
			uc.updateUser(id, w, r)
		case http.MethodDelete:
			uc.deleteUser(id, w)
		}
	}
}

func (uc userController) getUsers(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

func (uc userController) getUser(id int, w http.ResponseWriter) {
	user, err := models.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("User not found"))
		return
	}
	encodeResponseAsJSON(user, w)
}

func (uc userController) addUser(w http.ResponseWriter, r *http.Request) {
	user, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse user object in request"))
		return
	}
	user, err = models.AddUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(user, w)
}

func (uc userController) updateUser(id int, w http.ResponseWriter, r *http.Request) {
	user, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse user object in request"))
		return
	}
	user, err = models.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(user, w)
}

func (uc userController) deleteUser(id int, w http.ResponseWriter) {
	err := models.RemoveUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc userController) parseRequest(r *http.Request) (models.User, error) {
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Constructor function
func newUserController() *userController {
	// Creates a new userController and returns it's address. Go will keep track of the address even though this was declared locally here
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
