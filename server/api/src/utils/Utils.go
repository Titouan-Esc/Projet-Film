package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func ConsumApi(url string) (string, error) {
	var ouai string

	godotenv.Load()
	apiKey := os.Getenv("API_KEY")

	response, err := http.Get(url + apiKey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return ouai, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return ouai, err
	}

	return string(responseData), nil
}
