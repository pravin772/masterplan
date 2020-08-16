package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pravin772/mp-api/masterplan/model"
)

const (
	format = "2018-11-06"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m model.Masterplan
	b, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(b, &m)
	err = m.InsertMasterplan()
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
