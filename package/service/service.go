package service

import (
	// "mediumuz/model"
	"github.com/mrboburs/Norbekov/model"
	"github.com/mrboburs/Norbekov/package/repository"
	"github.com/mrboburs/Norbekov/util/logrus"
	"mime/multipart"
	// "mediumuz/util/logrus"
	// "mime/multipart"
)

type Home interface {
	CreateHomePost(post model.HomePost, logrus *logrus.Logger) (int, error)
	UpdateHomeImage(homeID int, filePath string, logrus *logrus.Logger) (int64, error)
	UploadImage(file multipart.File, header *multipart.FileHeader, logrus *logrus.Logger) (string, error)
	UpdateHome(id int, home model.HomePost, logrus *logrus.Logger) (int64, error)
	GetHomeById(id string, logrus *logrus.Logger) (model.HomeFull, error)
	DeleteHome(id string, logrus *logrus.Logger) error
	GetAllHome(logrus *logrus.Logger) (home []model.HomeFull, err error)
}
type News interface {
	CreateNewsPost(post model.NewsPost, logrus *logrus.Logger) (int, error)
	UpdateNewsImage(ID int, filePath string, logrus *logrus.Logger) (int64, error)
	UpdateNews(Id int, post model.NewsPost, logrus *logrus.Logger) (int64, error)
	GetNewsById(id string, logrus *logrus.Logger) (model.NewsFull, error)
	DeleteNews(id string, logrus *logrus.Logger) error
	GetAllNews(logrus *logrus.Logger) (array []model.NewsFull, err error)
	// CreateNewsPost(post model.NewsPost, logrus *logrus.Logger) (int, error)
}
type Services interface {
	CreateServicePost(post model.ServicePost, logrus *logrus.Logger) (int, error)
	UpdateServiceImage(ID int, filePath string, logrus *logrus.Logger) (int64, error)
	UpdateService(Id int, post model.ServicePost, logrus *logrus.Logger) (int64, error)
	GetServiceById(id string, logrus *logrus.Logger) (model.ServiceFull, error)
	DeleteService(id string, logrus *logrus.Logger) error
	GetAllService(logrus *logrus.Logger) (array []model.ServiceFull, err error)
}
type Table interface {
	CreateTablePost(post model.TablePost, logrus *logrus.Logger) (int, error)
	UpdateTableImage(ID int, filePath string, logrus *logrus.Logger) (int64, error)
	UpdateTable(Id int, post model.TablePost, logrus *logrus.Logger) (int64, error)
	GetTableById(id string, logrus *logrus.Logger) (model.TableFull, error)
	DeleteTable(id string, logrus *logrus.Logger) error
	GetAllTable(logrus *logrus.Logger) (array []model.TableFull, err error)
	GetAllCourse(logrus *logrus.Logger) (array []model.CourseFull1, err error)
	CreateCoursePost(post model.CourseFull, logrus *logrus.Logger) (int, error)
	DeleteCourse(id string, logrus *logrus.Logger) error
}
type Contact interface {
	CreateContactPost(post model.Contact, logrus *logrus.Logger) (int, error)
	GetAllContact(logrus *logrus.Logger) (contact []model.ContactFull, err error)
}
type Admin interface {
	CreateAdmin(adminData model.Admin, logrus *logrus.Logger) (int, error)
	GetAdmin(user_name, password string, logrus *logrus.Logger) (model.Admin, error)
	CheckAdmin(id int, logrus *logrus.Logger) (bool, error)
	DeleteAdmin(id string, logrus *logrus.Logger) error
	GenerateToken(user_name, password string, logrus *logrus.Logger) (string, error)
	ParseToken(accessToken string) (int, error)
}
type Service struct {
	Home
	News
	Services
	Table
	Contact
	Admin
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Home:     NewHomeService(repos.Home),
		News:     NewNewsService(repos.News),
		Services: NewServicesService(repos.Service),
		Contact:  NewContactService(repos.Contact),
		Table:    NewTableService(repos.Table),
		Admin:    NewAdminService(repos.Admin),
	}
}
