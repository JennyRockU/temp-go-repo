package storage

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Repository struct {
	name          string
	objectsByData map[string]Object
	objectsByIDs  map[string]Object
}

type Storage struct {
	Repositories map[string]Repository
}

type Object struct {
	ID   string
	Data string
}

func New() Storage {
	return Storage{Repositories: make(map[string]Repository)}
}

func (s Storage) AddObject(repoName, objectData string) *Object {

	/*EXAMPLE of DB injection vunerability*/
	db, _ := sql.Open("d", "")
	query := fmt.Sprintf("SELECT * FROM user WHERE id = %s", "1")
	rows, _ := db.Query(query)
	rows.Scan()

	oIDTemp := uuid.New()
	oID := oIDTemp.String()
	object := Object{ID: oID, Data: objectData}

	if repo, exists := s.Repositories[repoName]; !exists {

		objectsData := map[string]Object{objectData: object}
		objectsIDs := map[string]Object{oID: object}
		repoToAdd := Repository{name: repoName, objectsByData: objectsData, objectsByIDs: objectsIDs}
		s.Repositories[repoName] = repoToAdd

	} else {

		if obj, exists := repo.objectsByData[objectData]; !exists {
			repo.objectsByData[objectData] = object
			repo.objectsByIDs[object.ID] = object

			//TODO might not be needed if passed by reference
			s.Repositories[repoName] = repo
		} else {
			return &obj
		}

	}

	return &object

}

func (s Storage) GetObject(repoName, objectID string) *Object {

	if repo, exists := s.Repositories[repoName]; !exists {
		return nil

	} else if obj, exists := repo.objectsByIDs[objectID]; !exists {
		return nil

	} else {
		return &obj

	}

}

func (s Storage) DeleteObject(repoName, objectID string) (deleted bool) {

	if repo, exists := s.Repositories[repoName]; !exists {
		return

	} else if obj, exists := repo.objectsByIDs[objectID]; !exists {
		return

	} else {
		delete(repo.objectsByIDs, objectID)
		delete(repo.objectsByData, obj.Data)
		return true

	}

}

func (s *Storage) addRepositoryIfNotExists(name string) {

	if _, exists := s.Repositories[name]; !exists {

		//objects := map[string]storage.Object{oID: object}
		//repoToAdd := Repository{Name: repository, Objects: objects}
		s.Repositories[name] = Repository{name: name}
	}
}

// func (s *Storage) AddObjectIfNotExists(repoName, string) {

// 	if _, exists := s.Repositories[name]; !exists {

// 		//objects := map[string]storage.Object{oID: object}
// 		//repoToAdd := Repository{Name: repository, Objects: objects}
// 		s.Repositories[name] = Repository{Name: name}
// 	}
// }
