package main

import (
	"fmt"
	"net/http"

	chi "web-service/api/router"
)

func main() {
	fmt.Println("Listen to port 4000")
	http.ListenAndServe(":4000", chi.ChiRouter().InitRouter())
}
