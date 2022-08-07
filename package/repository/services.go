package repository

import (
	"fmt"
	"github.com/mrboburs/Norbekov/model"
	"github.com/mrboburs/Norbekov/util/logrus"
	"time"

	"github.com/jmoiron/sqlx"
	// "github.com/lib/pq"
)

type ServicePostDB struct {
	db *sqlx.DB
}

func NewServicesPostDB(db *sqlx.DB) *ServicePostDB {
	return &ServicePostDB{db: db}
}
func (repo *ServicePostDB) GetAllService(logrus *logrus.Logger) (array []model.ServiceFull, err error) {
	rowsRs, err := repo.db.Query("SELECT id,post_title,post_img_path,post_img_url, post_body, post_date ,price ,post_title_ru,post_body_ru FROM services")

	if err != nil {
		logrus.Infof("ERROR: not selecting data from sql %s", err.Error())
		// http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return array, err
	}

	Array := []model.ServiceFull{}
	defer rowsRs.Close()

	for rowsRs.Next() {
		snb := model.ServiceFull{}
		err = rowsRs.Scan(&snb.ID, &snb.PostTitle, &snb.PostImgPath, &snb.PostImgUrl, &snb.PostBody, &snb.PostDate, &snb.Price, &snb.PostTitleRu, &snb.PostBodyRu)
		if err != nil {
			logrus.Infof("ERROR: not scanning data from sql %s", err.Error())
			// log.Println(err)
			// http.Error(w, http.StatusText(500), 500)
			return array, err
		}
		Array = append(Array, snb)
	}

	if err = rowsRs.Err(); err != nil {

		return Array, err
	}
	return Array, nil
}
func (repo *ServicePostDB) CreateServicePost(post model.ServicePost, logrus *logrus.Logger) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (post_title ,post_img_url,   post_body,price,post_title_ru,post_body_ru ) VALUES ($1, $2, $3,$4,$5,$6)  RETURNING id", services)

	row := repo.db.QueryRow(query, post.PostTitle, post.PostImgUrl, post.PostBody, post.Price, post.PostTitleRu, post.PostBodyRu)

	if err := row.Scan(&id); err != nil {
		logrus.Infof("ERROR:PSQL Insert error %s", err.Error())
		return 0, err
	}
	logrus.Info("DONE: INSERTED Data PSQL")
	return id, nil
}

func (repo *ServicePostDB) UpdateServiceImage(ID int, filePath string, logrus *logrus.Logger) (int64, error) {
	tm := time.Now()
	query := fmt.Sprintf("UPDATE %s  SET post_img_path = $1,updated_at = $2    WHERE id = $3 RETURNING id", services)

	rows, err := repo.db.Exec(query, filePath, tm, ID)

	if err != nil {
		logrus.Errorf("ERROR: Update PostImage failed : %v", err)
		return 0, err
	}

	effectedRowsNum, err := rows.RowsAffected()

	if err != nil {
		logrus.Errorf("ERROR: Update Post Image effectedRowsNum : %v", err)
		return 0, err
	}
	logrus.Info("DONE:Update  image")
	return effectedRowsNum, nil

}

func (repo *ServicePostDB) UpdateService(Id int, post model.ServicePost, logrus *logrus.Logger) (int64, error) {
	tm := time.Now()
	query := fmt.Sprintf("	UPDATE %s SET post_title =$1, post_img_url  = $2, post_body = $3,  updated_at=$4, price=$5 ,post_title_ru=$6,post_body_ru=$7 WHERE id = $8 RETURNING id", "services")
	rows, err := repo.db.Exec(query, post.PostTitle, post.PostImgUrl, post.PostBody, tm, post.Price, post.PostTitleRu, post.PostBodyRu, Id)

	if err != nil {
		logrus.Errorf("ERROR: Update home : %v", err)
		return 0, err
	}
	effectedRowsNum, err := rows.RowsAffected()
	if err != nil {
		logrus.Errorf("ERROR: Update Home effectedRowsNum failed : %v", err)
		return 0, err
	}
	logrus.Info("DONE:Update l")
	return effectedRowsNum, nil
}
func (repo *ServicePostDB) DeleteService(id string, logrus *logrus.Logger) error {

	_, err := repo.db.Exec("DELETE from services WHERE id = $1", id)
	if err != nil {
		logrus.Errorf("ERROR: Update service : %v", err)
		return err
	}
	return nil
}

func (repo *ServicePostDB) GetServiceById(id string, logrus *logrus.Logger) (model.ServiceFull, error) {

	var post model.ServiceFull
	query := fmt.Sprintf("SELECT  id, post_title, post_img_path,post_img_url, post_body, post_date,price ,post_title_ru,post_body_ru  FROM %s WHERE id=$1 ", services)
	err := repo.db.Get(&post, query, id)
	if err != nil {
		logrus.Errorf("ERROR: don't get users %s", err)
		return post, err
	}
	logrus.Info("DONE:get user data from psql")

	return post, err
}
