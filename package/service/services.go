package service

import (
	"github.com/mrboburs/Norbekov/model"
	"github.com/mrboburs/Norbekov/package/repository"
	"github.com/mrboburs/Norbekov/util/logrus"
)

type ServicesService struct {
	repo repository.Service
}

func NewServicesService(repo repository.Service) *ServicesService {
	return &ServicesService{repo: repo}
}

func (s *ServicesService) CreateServicePost(post model.ServicePost, logrus *logrus.Logger) (int, error) {
	return s.repo.CreateServicePost(post, logrus)
}

func (s *ServicesService) UpdateServiceImage(ID int, filePath string, logrus *logrus.Logger) (int64, error) {
	return s.repo.UpdateServiceImage(ID, filePath, logrus)
}
func (s *ServicesService) UpdateService(Id int, post model.ServicePost, logrus *logrus.Logger) (int64, error) {
	return s.repo.UpdateService(Id, post, logrus)
}
func (s *ServicesService) GetServiceById(id string, logrus *logrus.Logger) (model.ServiceFull, error) {
	return s.repo.GetServiceById(id, logrus)
}
func (s *ServicesService) DeleteService(id string, logrus *logrus.Logger) error {
	return s.repo.DeleteService(id, logrus)
}
func (s *ServicesService) GetAllService(logrus *logrus.Logger) (array []model.ServiceFull, err error) {
	return s.repo.GetAllService(logrus)
}
