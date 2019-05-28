package handlers

import (
	"net/http"
)

// TriggerPanicHandler trigger panic handler
func TriggerPanicHandler(w http.ResponseWriter, r *http.Request) {

	panic("Triggering a Panic!")
}
