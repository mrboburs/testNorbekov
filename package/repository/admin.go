package repository

import (
	"fmt"
	"github.com/mrboburs/Norbekov/model"
	// "github.com/mrboburs/Norbekov/model"
	"github.com/mrboburs/Norbekov/util/logrus"
	// "time"

	"github.com/jmoiron/sqlx"
)

type AdminDB struct {
	db *sqlx.DB
}

func NewAdminDB(db *sqlx.DB) *AdminDB {
	return &AdminDB{db: db}
}

func (repo *AdminDB) CreateAdmin(adminData model.Admin, logrus *logrus.Logger) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_name ,password ) VALUES ($1, $2)  RETURNING id", admin)

	row := repo.db.QueryRow(query, adminData.UserName, adminData.Passord)

	if err := row.Scan(&id); err != nil {
		logrus.Infof("ERROR:PSQL Insert failed %s", err.Error())
		return 0, err
	}
	logrus.Info("DONE: INSERTED Data PSQL")
	return id, nil
}

func (repo *AdminDB) GetAdmin(user_name, password string, logrus *logrus.Logger) (model.Admin, error) {
	var Admin model.Admin
	query := fmt.Sprintf("SELECT id FROM %s WHERE user_name=$1 AND password=$2", admin)
	err := repo.db.Get(&Admin, query, user_name, password)
	// logrus.Fatal(err)
	if err != nil {
		return Admin, err
	}

	return Admin, nil
}

func (repo *AdminDB) CheckAdmin(id int, logrus *logrus.Logger) (bool, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE id=$1", admin)
	err := repo.db.Get(&count, query, id)

	if err != nil {
		logrus.Infof("ERROR: query error: %s", err.Error())
		return false, err
	}

	return true, nil
}

func (repo *AdminDB) DeleteAdmin(id string, logrus *logrus.Logger) error {

	_, err := repo.db.Exec("DELETE from admin WHERE id = $1", id)
	if err != nil {
		logrus.Errorf("ERROR: deleting failed : %v", err)
		return err
	}
	return nil
}
