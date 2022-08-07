package repository

import (
	"fmt"
	"github.com/mrboburs/Norbekov/model"
	"github.com/mrboburs/Norbekov/util/logrus"
	"time"

	"github.com/jmoiron/sqlx"
	// "github.com/lib/pq"
	// "github.com/lib/pq"
	// "github.com/sirupsen/logrus"
)

type HomePostDB struct {
	db *sqlx.DB
}

func NewHomePostDB(db *sqlx.DB) *HomePostDB {
	return &HomePostDB{db: db}
}

func (repo *HomePostDB) GetAllHome(logrus *logrus.Logger) (array []model.HomeFull, err error) {
	rowsRs, err := repo.db.Query("SELECT id,post_title,post_title_ru,post_img_path,post_img_url, post_body,post_body_ru, post_date  FROM home")

	if err != nil {
		logrus.Infof("ERROR: not selecting data from sql %s", err.Error())
		// http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return array, err
	}

	Array := []model.HomeFull{}
	defer rowsRs.Close()

	for rowsRs.Next() {
		snb := model.HomeFull{}
		err = rowsRs.Scan(&snb.ID, &snb.PostTitle, &snb.PostTitleRu, &snb.PostImgPath, &snb.PostImgUrl, &snb.PostBody, &snb.PostBodyRu, &snb.PostDate)
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

func (repo *HomePostDB) GetHomeById(id string, logrus *logrus.Logger) (model.HomeFull, error) {

	var post model.HomeFull
	query := fmt.Sprintf("SELECT  id, post_title,post_title_ru, post_img_path,post_img_url, post_body,post_body_ru, post_date  FROM %s WHERE id=$1 ", home)
	err := repo.db.Get(&post, query, id)
	if err != nil {
		logrus.Errorf("ERROR: don't get users %s", err)
		return post, err
	}
	logrus.Info("DONE:get user data from psql")

	return post, err
}

func (repo *HomePostDB) CreateHomePost(post model.HomePost, logrus *logrus.Logger) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (post_title,post_title_ru,post_body_ru ,post_img_url,   post_body ) VALUES ($1, $2, $3,$4,$5)  RETURNING id", home)

	row := repo.db.QueryRow(query, post.PostTitle, post.PostTitleRu, post.PostBodyRu, post.PostImgUrl, post.PostBody)

	if err := row.Scan(&id); err != nil {
		logrus.Infof("ERROR:PSQL Insert error %s", err.Error())
		return 0, err
	}
	logrus.Info("DONE: INSERTED Data PSQL")
	return id, nil
}

func (repo *HomePostDB) UpdateHomeImage(id int, filePath string, logrus *logrus.Logger) (int64, error) {
	tm := time.Now()
	query := fmt.Sprintf("UPDATE %s  SET post_img_path = $1,updated_at = $2    WHERE id = $3 RETURNING id", home)

	rows, err := repo.db.Exec(query, filePath, tm, id)

	if err != nil {
		logrus.Errorf("ERROR: Update HomeImage failed : %v", err)
		return 0, err
	}

	effectedRowsNum, err := rows.RowsAffected()

	if err != nil {
		logrus.Errorf("ERROR: Update Post Image effectedRowsNum : %v", err)
		return 0, err
	}
	logrus.Info("DONE:Update Homr image")
	return effectedRowsNum, nil

}

func (repo *HomePostDB) UpdateHome(Id int, post model.HomePost, logrus *logrus.Logger) (int64, error) {
	tm := time.Now()
	query := fmt.Sprintf("	UPDATE %s SET post_title =$1, post_img_url  = $2, post_body = $3,  updated_at=$4,post_title_ru=$5,post_body_ru=$6 WHERE id = $7 RETURNING id", home)
	rows, err := repo.db.Exec(query, post.PostTitle, post.PostImgUrl, post.PostBody, tm, post.PostTitleRu, post.PostBodyRu, Id)

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
func (repo *HomePostDB) DeleteHome(id string, logrus *logrus.Logger) error {

	_, err := repo.db.Exec("DELETE from home WHERE id = $1", id)
	if err != nil {
		logrus.Errorf("ERROR: Update home : %v", err)
		return err
	}
	return nil
}
