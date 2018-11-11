package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"../model"
)

type ManagementType struct {
	data *model.Management
}

func (m *ManagementType) registerRoute() {
	http.HandleFunc("/manager/record", m.ManagerRecord)
	http.HandleFunc("/manager/show", m.ManagerShow)
}

func (m *ManagementType) ManagerRecord(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data := model.Management{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = model.RecordManager(&data)
	if err != nil {
		log.Println(err.Error())
		return
	}
	w.Write([]byte("Inserted successfully"))
	return
}

func (m *ManagementType) ManagerShow(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != os.Getenv("ADMIN_PASS") {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Sorry, forbidden"))
		return
	}

	data := &model.Management{}
	query := r.URL.Query().Get("reg")

	if len(query) < 1 {
		arr, err := model.ShowAllmanager()
		if err != nil {
			log.Println(err.Error())
			return
		}
		json.NewEncoder(w).Encode(arr)
	}

	data, err := model.Showmanager(query)
	if err != nil {
		log.Println(err.Error())
		return
	}
	json.NewEncoder(w).Encode(data)
}
