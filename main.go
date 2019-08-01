package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Theory Struct(Model)
type Theory struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Text    string `json:"text"`
	Picture string `json:"picture"`
}

var theoryArray []Theory

func getTheory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(theoryArray)
}
func main() {
	router := mux.NewRouter()
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/math_theory")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err := db.Query("Select id,title,text,picture from theory")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var theory Theory
		err = results.Scan(&theory.ID, &theory.Title, &theory.Text, &theory.Picture)

		if err != nil {
			panic(err.Error())
		}
		theoryArray = append(theoryArray, theory)

	}
	router.HandleFunc("/api/theories", getTheory).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

}
