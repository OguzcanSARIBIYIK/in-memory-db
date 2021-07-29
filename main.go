package main

import (
	"encoding/json"
	"fmt"
	"in-memory-db/indb"
	"net/http"
)

var DB *indb.InDB

type Response struct {
	Code int
	Data interface{}
}

func main() {
	DB = indb.Init()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		key := r.URL.Query().Get("key")
		if key == "" {
			rw.WriteHeader(http.StatusBadRequest)
			res, _ := json.Marshal(Response{Code: http.StatusBadRequest, Data: "Key is required."})
			rw.Write(res)
			return
		}
		value := r.URL.Query().Get("value")
		if value == "" {
			rw.WriteHeader(http.StatusBadRequest)
			res, _ := json.Marshal(Response{Code: http.StatusBadRequest, Data: "Value is required."})
			rw.Write(res)
			return
		}

		rw.Header().Set("Content-Type", "application/json")

		DB.Save(key, value)
		res, _ := json.Marshal(Response{Code: http.StatusOK, Data: "Data saved."})
		rw.Write(res)

	})

	http.HandleFunc("/all", func(rw http.ResponseWriter, r *http.Request) {
		data := DB.GetAll()

		rw.Header().Set("Content-Type", "application/json")
		res, _ := json.Marshal(Response{Code: http.StatusOK, Data: data})
		rw.Write(res)
	})

	fmt.Println("Listening :8889...")
	err := http.ListenAndServe(":8889", nil)

	if err != nil {
		panic(err)
	}

}
