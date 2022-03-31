package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()
	// Handle's 2nd param takes an http Handler interface. Our uc is already one because we defined ServeHTTP correctly
	http.Handle("/users", *uc)
	// To be able to handle both requests to /users and requests to sub routes from /users/ we have to define both
	http.Handle("/users/", *uc)
}

// io.Writer here is a parent interface of the http.ResponseWriter. This is because the response writer has the
// Write([]byte) (int, error) function which corresponds to the Write(p []byte) (n int, err error) function in the io Writer
func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
