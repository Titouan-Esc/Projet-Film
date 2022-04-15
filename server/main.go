package main

import (
	"fmt"
	"net/http"

	chi "web-service/api/router"
)

func main() {
	fmt.Println("Listen to port 4100")
	http.ListenAndServe(":4100", chi.ChiRouter().InitRouter())
}
