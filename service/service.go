package service

// import (
// 	"storage-api/storage"

// 	"github.com/google/uuid"
// )

// type Service struct {
// 	storage Storage
// }

// func New(storage Storage) *Service {
// 	return &Service{Storage: sstorage}

// }

// func (s *Service) getObject(repository, objectID string) {

// }

// func (s *Service) addObject(repository, data string) {

// 	oIDTemp := uuid.New()
// 	oID := oIDTemp.String()
// 	object := storage.Object{ID: oID, Data: data}

// 	s.storage.AddRepositoryIfNotExists(repository)

// 	if repo, exists := s.storage.Repositories[repository]; !exists {

// 		objects := map[string]storage.Object{oID: object}
// 		repoToAdd := storage.Repository{Name: repository, Objects: objects}
// 		s.storage.Repositories[repository] = repoToAdd
// 	} else {
// 		if _, exists := repo.Objects[data]; !exists {
// 			repo.Objects = append(repo.Objects, object)
// 			s.storage.Repositories[repository] = repo
// 		}

// 	}

// }
