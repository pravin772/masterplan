package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/pravin772/mp-api/masterplan/controller"
	"github.com/pravin772/mp-api/masterplan/model"
)

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", controller.HelloHandler).Methods("GET")
	myRouter.HandleFunc("/addData", controller.AddData).Methods("POST")
	// myRouter.HandleFunc("/deleteData", controller.DeleteData).Methods("DELETE")
	// myRouter.HandleFunc("/getData", controller.GetData).Methods("GET")
	// myRouter.HandleFunc("/getAll", controller.GetAllData).Methods("GET")
	// myRouter.HandleFunc("/dcsv", csv_generator.DownloadCSV).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	model.DB, model.Err = gorm.Open("mysql", "user1:user1@(localhost)/masterplan?charset=utf8&parseTime=True&loc=Local")
	if model.Err != nil {
		fmt.Println(model.Err.Error())
		panic("Failed to open database")
	}
	defer model.DB.Close()
	model.InitMigration()
	fmt.Println("Server running on localhost:8000")
	handleRequest()
}
