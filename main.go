package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/pravin772/mp-api/masterplan/controller"
	"github.com/pravin772/mp-api/masterplan/csv_generator"
	"github.com/pravin772/mp-api/masterplan/model"
)

func setContent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", controller.HelloHandler).Methods("GET")
	myRouter.Use(setContent)
	myRouter.HandleFunc("/addData", controller.AddData).Methods("POST")
	myRouter.HandleFunc("/getAll", controller.GetAllActivities).Methods("GET")

	// DownloadCSV api sorted by default SrNo(WBS)
	myRouter.HandleFunc("/dcsv", csv_generator.DownloadCSV).Methods("GET")
	// DownloadCSV by StartDate in ascending order
	myRouter.HandleFunc("/dcsvbystartdate", csv_generator.GetAllActivitiesByStartDate).Methods("GET")

	s := &http.Server{
		Addr:           ":8000",
		Handler:        myRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func loadEnvironment() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	loadEnvironment()
	model.DB, model.Err = gorm.Open("mysql", os.Getenv("DB_URL"))
	if model.Err != nil {
		log.Println(model.Err.Error())
		panic("Failed to open database")
	}
	defer model.DB.Close()
	model.InitMigration()
	log.Println("Server running on localhost:8000")
	handleRequest()
}
