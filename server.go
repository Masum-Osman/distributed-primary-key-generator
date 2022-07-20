package main

import (
	sf "dpkg/snowflake"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type KeyParams struct {
	DataCenterID uint64 `json:"datacenterID"`
	MachineID    uint64 `json:"machineID"`
}

type UniqueID struct {
	Key uint64 `json:"key"`
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

	var res UniqueID
	res.Key = sf.Snowflake(params.DataCenterID, params.MachineID)
	output, err := json.Marshal(res)
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
