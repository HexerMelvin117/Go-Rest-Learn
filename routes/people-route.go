package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hexermelvin117/goback/models"
)

var people []models.Person

func fillPeople() {
	people = []models.Person{
		{Id: 1, Name: "Henry", Lastname: "Oxford", Age: 22},
		{Id: 2, Name: "John", Lastname: "Klein", Age: 43},
		{Id: 3, Name: "Kerry", Lastname: "Knur", Age: 23},
	}
}

func ReturnAllPeople(w http.ResponseWriter, r *http.Request) {
	// This is a test
	fillPeople()
	fmt.Println("Endpoint Hit: returnAllPeople")
	json.NewEncoder(w).Encode(people)
}
