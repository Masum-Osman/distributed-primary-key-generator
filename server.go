package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type KeyParams struct {
	DataCenterID int32 `json:"datacenterID"`
	MachineID    int32 `json:"machineID"`
}

func logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		next(w, r)
	}
}

func generate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var params KeyParams
	err = json.Unmarshal(body, &params)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(params)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)

}

func main() {
	http.HandleFunc("/generate-key", logging(generate))
	http.ListenAndServe(":8083", nil)
}
