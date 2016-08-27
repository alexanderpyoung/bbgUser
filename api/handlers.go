package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

type UserParams struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	FriendlyName string `json:"friend"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	p := UserParams{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &p)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	upValue := reflect.ValueOf(&p).Elem()
	for i := 0; i < upValue.NumField(); i++ {
		if upValue.Field(i).Type() == reflect.TypeOf("") && upValue.Field(i).String() == "" {
			fmt.Printf("Error: incomplete json message")
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func viewAccountHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
}
