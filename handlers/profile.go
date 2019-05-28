package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ProfileHandler profile handler
func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]
	w.Write([]byte(username))
}
