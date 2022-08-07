package service

import (
	"github.com/mrboburs/Norbekov/model"
	"github.com/mrboburs/Norbekov/package/repository"
	"github.com/mrboburs/Norbekov/util/logrus"
)

type TableService struct {
	repo repository.Table
}

func NewTableService(repo repository.Table) *TableService {
	return &TableService{repo: repo}
}

func (s *TableService) CreateCoursePost(post model.CourseFull, logrus *logrus.Logger) (int, error) {

	return s.repo.CreateCoursePost(post, logrus)
}
func (s *TableService) CreateTablePost(post model.TablePost, logrus *logrus.Logger) (int, error) {

	return s.repo.CreateTablePost(post, logrus)
}
func (s *TableService) UpdateTableImage(ID int, filePath string, logrus *logrus.Logger) (int64, error) {
	return s.repo.UpdateTableImage(ID, filePath, logrus)
}

func (s *TableService) UpdateTable(Id int, post model.TablePost, logrus *logrus.Logger) (int64, error) {
	return s.repo.UpdateTable(Id, post, logrus)
}

func (s *TableService) GetTableById(id string, logrus *logrus.Logger) (model.TableFull, error) {
	return s.repo.GetTableById(id, logrus)
}
func (s *TableService) DeleteCourse(id string, logrus *logrus.Logger) error {
	return s.repo.DeleteTable(id, logrus)
}
func (s *TableService) GetAllTable(logrus *logrus.Logger) (array []model.TableFull, err error) {
	return s.repo.GetAllTable(logrus)
}
func (s *TableService) DeleteTable(id string, logrus *logrus.Logger) error {
	return s.repo.DeleteTable(id, logrus)
}
func (s *TableService) GetAllCourse(logrus *logrus.Logger) (array []model.CourseFull1, err error) {
	return s.repo.GetAllCourse(logrus)
}
