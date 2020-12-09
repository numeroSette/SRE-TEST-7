package getrandomnumbernative

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// randomInt returns a string with random numbers
func randomInt() string {

	// Reference
	// https://flaviocopes.com/go-random/
	// https://www.random.org/sequences/?min=1&max=52&col=1&format=plain&rnd=new

	// Use Unix date to seed
	rand.Seed(time.Now().UnixNano())

	// Generate 52 numbers between 1 and 60
	var array = make([]int, 52)

	for i := 0; i < 52; i++ {
		array[i] = 1 + rand.Intn(60-1)
	}

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), ""), "[]")
}

// RandomNumberResponse is a struct for response
type RandomNumberResponse struct {
	RandomNumber string `json:"random_number"`
}

// GetRandomNumberNative is a function to return a random number
func GetRandomNumberNative(response http.ResponseWriter, request *http.Request) {

	out := &RandomNumberResponse{
		RandomNumber: randomInt(),
	}

	response.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(response).Encode(out)
	if err != nil {
		log.Printf("failed to encode json to HTTP response: %v", err)
	}

}
