package csv_generator

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sort"

	"github.com/pravin772/mp-api/masterplan/model"
)

const (
	formatter = "01-02-2006"
)

func DownloadCSV(w http.ResponseWriter, r *http.Request) {
	data, err := model.GetAllActivities()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(err.Error())
		return
	}

	buffer := &bytes.Buffer{}
	writer := csv.NewWriter(buffer)
	row := []string{"SrNo", "Activity", "Start Date(MM-DD-YYYY)", "End Date"}
	err = writer.Write(row)
	if err := writer.Error(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, field := range data {
		row := []string{field.SrNo, field.Activity, field.StartDate.Format(formatter), field.EndDate.Format(formatter)}
		err := writer.Write(row)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Setting the content type header to text/csv because our middleware bydefault sets it to json
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=AllActivities.csv")
	//w.Write(buffer.Bytes())
	// This will stream the file
	io.Copy(w, buffer)
	log.Println(r.RequestURI, " served")
}

// By start dates
func GetAllActivitiesByStartDate(w http.ResponseWriter, r *http.Request) {
	data, err := model.GetAllActivities()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(err.Error())
		return
	}

	buffer := &bytes.Buffer{}
	writer := csv.NewWriter(buffer)
	row := []string{"Start Date(MM-DD-YYYY)", "End Date", "SrNo", "Activity"}
	err = writer.Write(row)
	if err := writer.Error(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sort.Sort(byStartDate(data))
	log.Println("Data Sorted")

	for _, field := range data {
		row := []string{field.StartDate.Format(formatter), field.EndDate.Format(formatter), field.SrNo, field.Activity}
		err := writer.Write(row)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// setting the content type header to text/csv because our middleware bydefault sets it to json
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=AllActivitiesByStartDate.csv")
	//w.Write(buffer.Bytes())
	io.Copy(w, buffer)
	log.Println(r.RequestURI, " served")
}

type byStartDate []*model.Activity

func (a byStartDate) Len() int           { return len(a) }
func (a byStartDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byStartDate) Less(i, j int) bool { return a[i].StartDate.Before(a[j].StartDate) }
