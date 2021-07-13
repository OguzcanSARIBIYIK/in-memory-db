package main

import (
	"encoding/json"
	"fmt"
	"in-memory-db/indb"
	"io/ioutil"
	"log"
	"net/http"
)

var DB *indb.InDB

type Add struct {
	Key   *string
	Value *interface{}
}

type Response struct {
	Code int
	Data interface{}
}

func main() {
	DB = indb.Init()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		var req Add

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(reqBody, &req)

		rw.Header().Set("Content-Type", "application/json")

		if req.Value != nil && req.Key != nil {
			fmt.Println(*req.Value)
			DB.Save(*req.Key, *req.Value)
			res, _ := json.Marshal(Response{Code: http.StatusOK, Data: "Veri Kaydedildi."})
			rw.Write(res)
		} else {
			rw.WriteHeader(http.StatusBadRequest)
			res, _ := json.Marshal(Response{Code: http.StatusBadRequest, Data: "Veri Kaydedilemedi."})
			rw.Write(res)
		}

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

	/* db.Save("Age", "23")
	db.Save("Age", "31")
	age, err := db.Get("Age")
	fmt.Println(age, err)
	//db.Delete("Age")
	db.Save("Name", "Oğuzcan")
	name, err := db.Get("Name")
	fmt.Println(name, err)
	allData := db.GetAll()

	for k, v := range allData {
		fmt.Printf("Key : %v - Value : %v \n", k, v)
	} */
}
