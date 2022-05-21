package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("app.env")

	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Get("/", handler)
	r.Get("/test", testHandler)

	InitDb()
	http.ListenAndServe(":8300", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var dum dummy
	num, _ := strconv.Atoi(r.URL.Query().Get("id"))
	err := Db.Get(&dum, "SELECT * FROM dummytable WHERE id = $1", num)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	j, _ := json.Marshal(dum)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

type dummy struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
