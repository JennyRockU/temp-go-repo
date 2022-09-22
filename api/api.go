package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	store "storage-api/storage"

	"github.com/gorilla/mux"
)

type CreatedResponse struct {
	ObjectID string `json:"oid"`
	Size     int    `json:"size"`
}

type GetResponse struct {
	Object string `json:"object"`
}

func ServerHTTP() {
	router := mux.NewRouter()

	router.Methods(http.MethodPut).
		Path("/data/{repository}").
		HandlerFunc(addObject)

	router.Methods(http.MethodGet).
		Path("/data/{repository}/{objectID}").
		HandlerFunc(getObject)

}

type storageHandler interface {
	AddObject(repoName, objectData string) *store.Object
	GetObject(repoName, objectID string) *store.Object
	DeleteObject(repoName, objectID string) (deleted bool)
}

var storage storageHandler

func init() {
	storage = store.New()
}

/*
Status: 201 Created
{
  "oid": "2845f5a412dbdfacf95193f296dd0f5b2a16920da5a7ffa4c5832f223b03de96",
  "size": 1234
}
*/
func addObject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	repo := vars["repository"]

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := string(bodyBytes)
	fmt.Printf("storing the data... %v for repo %v", data, repo)
	obj := storage.AddObject(repo, data)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreatedResponse{ObjectID: obj.ID, Size: len(obj.Data)})

}

/*
Status: 200 OK
{object data}
*/
func getObject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	repo, objectID := vars["repository"], vars["objectID"]

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := string(bodyBytes)
	fmt.Printf("getting the data... %v for repo %v, objID", data, repo, objectID)

	obj := storage.GetObject(repo, objectID)
	if obj == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(GetResponse{data})

}

/*
Status: 200 OK
*/
func deleteObject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	repo, objectID := vars["repository"], vars["objectID"]

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := string(bodyBytes)
	fmt.Printf("deleting the data... %v for repo %v, objID", data, repo, objectID)
	ok := storage.DeleteObject(repo, objectID)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

}
