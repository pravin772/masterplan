package csv_generator

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pravin772/mp-api/masterplan/model"
)

const (
	formatter = "2018-11-06"
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
	row := []string{"SrNo", "Activity", "Start Date", "End Date"}
	err = writer.Write(row)
	if err := writer.Error(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, field := range data {
		row := []string{field.SrNo, field.Activity, field.StartDate.String(), field.EndDate.String()}
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
	w.Header().Set("Content-Disposition", "attachment;filename=AllActivities.csv")
	//w.Write(buffer.Bytes())
	io.Copy(w, buffer)
}

//by start dates
func GetAllActivitiesByStartDate(w http.ResponseWriter, r *http.Request) {
	data, err := model.GetAllActivities()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(err.Error())
		return
	}

}
