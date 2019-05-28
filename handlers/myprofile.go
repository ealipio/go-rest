package handlers

import "net/http"

// MyProfileHandler my profile handler
func MyProfileHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("profile"))

}
