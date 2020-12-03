package getrandomnumber

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const (
	RandomNumberGetSource = "https://www.random.org/sequences/?min=1&max=52&col=1&format=plain&rnd=new"
)

var (
	RandomNumberRegexMatch = regexp.MustCompile("\n")
)

type RandomNumberResponse struct {
	RandomNumber string `json:"random_number"`
}

// GetRandomNumber is a function to return a random number
func GetRandomNumber(response http.ResponseWriter, request *http.Request) {

	sourceResponse, err := http.Get(RandomNumberGetSource)
	if err != nil {
		log.Println(err)
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	defer sourceResponse.Body.Close()

	sourceBody, err := ioutil.ReadAll(sourceResponse.Body)
	if err != nil {
		log.Println(err)
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	out := &RandomNumberResponse{
		RandomNumber: RandomNumberRegexMatch.ReplaceAllString(string(sourceBody), ""),
	}

	response.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(response).Encode(out)
	if err != nil {
		log.Printf("failed to encode json to HTTP response: %v", err)
	}

}
