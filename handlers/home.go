package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler home handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("test")

}
