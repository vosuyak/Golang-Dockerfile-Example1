package main

import (
	"fmt"
	"net/http"
)

// handler to display a successful connection with docker
func connectDocker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "docker file running! ...")
}

func main() {
	http.HandleFunc("/", connectDocker)
	http.ListenAndServe(":8081", nil)
}
