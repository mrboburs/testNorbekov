package repository

import (
	// "norbekov/model"
	// "norbekov/util/logrus"

	// "github.com/go-redis/redis"
	"github.com/mrboburs/Norbekov/model"
	"github.com/mrboburs/Norbekov/util/logrus"

	"github.com/jmoiron/sqlx"
)

type Home interface {
	CreateHomePost(post model.HomePost, logrus *logrus.Logger) (int, error)
	UpdateHomeImage(homeID int, filePath string, logrus *logrus.Logger) (int64, error)
	UpdateHome(homeId int, home model.HomePost, logrus *logrus.Logger) (int64, error)
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
}
type Service interface {
	GetAllService(logrus *logrus.Logger) (array []model.ServiceFull, err error)
	CreateServicePost(post model.ServicePost, logrus *logrus.Logger) (int, error)
	UpdateServiceImage(ID int, filePath string, logrus *logrus.Logger) (int64, error)
	UpdateService(Id int, post model.ServicePost, logrus *logrus.Logger) (int64, error)
	GetServiceById(id string, logrus *logrus.Logger) (model.ServiceFull, error)
	DeleteService(id string, logrus *logrus.Logger) error
}
type Table interface {
	GetAllTable(logrus *logrus.Logger) (array []model.TableFull, err error)
	UpdateTableImage(ID int, filePath string, logrus *logrus.Logger) (int64, error)
	UpdateTable(Id int, post model.TablePost, logrus *logrus.Logger) (int64, error)

	GetAllCourse(logrus *logrus.Logger) (array []model.CourseFull1, err error)
	CreateCoursePost(post model.CourseFull, logrus *logrus.Logger) (int, error)
	DeleteCourse(id string, logrus *logrus.Logger) error
	CreateTablePost(post model.TablePost, logrus *logrus.Logger) (int, error)
	DeleteTable(id string, logrus *logrus.Logger) error
	GetTableById(id string, logrus *logrus.Logger) (model.TableFull, error)
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
}
type Repository struct {
	Home
	News
	Service
	Table
	Contact
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{

		Home:    NewHomePostDB(db),
		News:    NewNewsPostDB(db),
		Service: NewServicesPostDB(db),
		Table:   NewTablesPostDB(db),
		Contact: NewContactPostDB(db),
		Admin:   NewAdminDB(db),
	}
}
