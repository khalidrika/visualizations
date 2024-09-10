package groupie

import (
	"encoding/json"
	"log"
	"net/http"
	groupie "groupie/data"
)

func HomeApi(url string, data []groupie.Band) []groupie.Band {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
