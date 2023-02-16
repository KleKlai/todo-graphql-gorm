package service

import (
	repository "github.com/kleklai/todoAppv1/repository"
)

type Service struct {
	repoService repository.Repository
}

func NewService(repoService repository.Repository) *Service {
	return &Service{
		repoService: repoService,
	}
}

// func (e *errorString) Error() string {
// 	return e.s
// }
