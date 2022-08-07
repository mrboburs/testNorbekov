package service

import (
	// "norbekov/model"
	"github.com/mrboburs/Norbekov/model"
	"github.com/mrboburs/Norbekov/package/repository"
	"github.com/mrboburs/Norbekov/util/logrus"
	// "norbekov/util/logrus"
)

type NewsService struct {
	repo repository.News
}

func NewNewsService(repo repository.News) *NewsService {
	return &NewsService{repo: repo}
}

func (s *NewsService) CreateNewsPost(post model.NewsPost, logrus *logrus.Logger) (int, error) {
	homeId, err := s.repo.CreateNewsPost(post, logrus)
	if err != nil {
		return 0, err
	}
	return homeId, err
}
func (s *NewsService) UpdateNewsImage(ID int, filePath string, logrus *logrus.Logger) (int64, error) {
	return s.repo.UpdateNewsImage(ID, filePath, logrus)
}
func (s *NewsService) UpdateNews(id int, post model.NewsPost, logrus *logrus.Logger) (int64, error) {
	return s.repo.UpdateNews(id, post, logrus)
}

func (s *NewsService) GetNewsById(id string, logrus *logrus.Logger) (model.NewsFull, error) {
	return s.repo.GetNewsById(id, logrus)
}
func (s *NewsService) DeleteNews(id string, logrus *logrus.Logger) error {
	return s.repo.DeleteNews(id, logrus)
}
func (s *NewsService) GetAllNews(logrus *logrus.Logger) (array []model.NewsFull, err error) {
	return s.repo.GetAllNews(logrus)
}
