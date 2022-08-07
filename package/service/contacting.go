package service

import (
	"github.com/mrboburs/Norbekov/model"
	"github.com/mrboburs/Norbekov/package/repository"
	"github.com/mrboburs/Norbekov/util/logrus"
)

type ContactService struct {
	repo repository.Contact
}

func NewContactService(repo repository.Contact) *ContactService {
	return &ContactService{repo: repo}
}

func (s *ContactService) CreateContactPost(post model.Contact, logrus *logrus.Logger) (int, error) {
	return s.repo.CreateContactPost(post, logrus)
}
func (s *ContactService) GetAllContact(logrus *logrus.Logger) (contact []model.ContactFull, err error) {
	return s.repo.GetAllContact(logrus)
}
