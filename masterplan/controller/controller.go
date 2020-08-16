package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pravin772/mp-api/masterplan/model"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func AddData(w http.ResponseWriter, r *http.Request) {
	var m model.Activity
	b, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(b, &m)
	err = m.InsertActivity()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetAllActivities(w http.ResponseWriter, r *http.Request) {
	data, err := model.GetAllActivities()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
