package main

import (
	"fmt"
	"net/http"
	"os"

	chi "web-service/api/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4200"
	}
	fmt.Printf("Listen to port %s\n", port)
	http.ListenAndServe(":"+port, chi.ChiRouter().InitRouter())
}
