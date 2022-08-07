package repository

import (
	"fmt"
	"github.com/mrboburs/Norbekov/model"
	"github.com/mrboburs/Norbekov/util/logrus"
	"time"

	"github.com/jmoiron/sqlx"
	// "github.com/lib/pq"
)

type TablePostDB struct {
	db *sqlx.DB
}

func NewTablesPostDB(db *sqlx.DB) *TablePostDB {
	return &TablePostDB{db: db}
}
func (repo *TablePostDB) GetAllCourse(logrus *logrus.Logger) (array []model.CourseFull1, err error) {
	rowsRs, err := repo.db.Query("SELECT  id, title ,body ,price,duration,term,format,created_at ,title_ru,body_ru FROM course")

	if err != nil {
		logrus.Infof("ERROR: not selecting data from sql %s", err.Error())
		// http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return array, err
	}

	Array := []model.CourseFull1{}
	defer rowsRs.Close()

	for rowsRs.Next() {
		snb := model.CourseFull1{}
		err = rowsRs.Scan(&snb.ID, &snb.Title, &snb.Body, &snb.Price, &snb.Duration, &snb.Term, &snb.Format, &snb.Date, &snb.TitleRu, &snb.BodyRu)
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
func (repo *TablePostDB) GetAllTable(logrus *logrus.Logger) (array []model.TableFull, err error) {
	rowsRs, err := repo.db.Query("SELECT id,post_title,post_img_path,post_img_url, post_body, post_date,price,duration ,post_title_ru,post_body_ru ,date,format FROM tables")

	if err != nil {
		logrus.Infof("ERROR: not selecting data from sql %s", err.Error())
		// http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return array, err
	}

	Array := []model.TableFull{}
	defer rowsRs.Close()

	for rowsRs.Next() {
		snb := model.TableFull{}
		err = rowsRs.Scan(&snb.ID, &snb.PostTitle, &snb.PostImgPath, &snb.PostImgUrl, &snb.PostBody, &snb.PostDate, &snb.Price, &snb.Duration, &snb.PostTitleRu, &snb.PostBodyRu, &snb.Date, &snb.Format)
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
func (repo *TablePostDB) CreateTablePost(post model.TablePost, logrus *logrus.Logger) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (post_title ,post_img_url,   post_body ,price,duration,post_title_ru,post_body_ru,date,format) VALUES ($1, $2, $3,$4,$5,$6,$7,$8,$9)  RETURNING id", tables)

	row := repo.db.QueryRow(query, post.PostTitle, post.PostImgUrl, post.PostBody, post.Price, post.Duration, post.PostTitleRu, post.PostBodyRu, post.Date, post.Format)

	if err := row.Scan(&id); err != nil {
		logrus.Infof("ERROR:PSQL Insert error %s", err.Error())
		return 0, err
	}
	logrus.Info("DONE: INSERTED Data PSQL")
	return id, nil
}

func (repo *TablePostDB) CreateCoursePost(post model.CourseFull, logrus *logrus.Logger) (int, error) {
	var id int
	tm := time.Now()
	query := fmt.Sprintf("INSERT INTO %s (title ,body ,price,duration,term,format,created_at,title_ru,body_ru) VALUES ($1, $2, $3,$4,$5,$6,$7,$8,$9)  RETURNING id", "course")

	row := repo.db.QueryRow(query, post.Title, post.Body, post.Price, post.Duration, post.Term, post.Format, tm, post.TitleRu, post.BodyRu)

	if err := row.Scan(&id); err != nil {
		logrus.Infof("ERROR:PSQL Insert error %s", err.Error())
		return 0, err
	}
	logrus.Info("DONE: INSERTED Data PSQL")
	return id, nil
}

func (repo *TablePostDB) UpdateTableImage(ID int, filePath string, logrus *logrus.Logger) (int64, error) {
	tm := time.Now()
	query := fmt.Sprintf("UPDATE %s  SET post_img_path = $1,updated_at = $2    WHERE id = $3 RETURNING id", tables)

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

func (repo *TablePostDB) UpdateTable(Id int, post model.TablePost, logrus *logrus.Logger) (int64, error) {
	tm := time.Now()
	query := fmt.Sprintf("	UPDATE %s SET post_title =$1, post_img_url  = $2, post_body = $3,  updated_at=$4,price=$5,duration=$6,post_title_ru=$7,post_body_ru=$8 ,date=$9,format=$10 WHERE id = $11 RETURNING id", tables)
	rows, err := repo.db.Exec(query, post.PostTitle, post.PostImgUrl, post.PostBody, tm, post.Price, post.Duration, post.PostTitleRu, post.PostBodyRu, post.Date, post.Format, Id)

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

func (repo *TablePostDB) DeleteTable(id string, logrus *logrus.Logger) error {

	_, err := repo.db.Exec("DELETE from tables WHERE id = $1", id)
	if err != nil {
		logrus.Errorf("ERROR: Deleting table : %v", err)
		return err
	}
	return nil
}
func (repo *TablePostDB) DeleteCourse(id string, logrus *logrus.Logger) error {

	_, err := repo.db.Exec("DELETE from course WHERE id = $1", id)
	if err != nil {
		logrus.Errorf("ERROR: Deleting table : %v", err)
		return err
	}
	return nil
}

func (repo *TablePostDB) GetTableById(id string, logrus *logrus.Logger) (model.TableFull, error) {

	var post model.TableFull
	query := fmt.Sprintf("SELECT  id, post_title, post_img_path,post_img_url, post_body, post_date,price,duration ,post_title_ru,post_body_ru,date,format FROM %s WHERE id=$1 ", tables)
	err := repo.db.Get(&post, query, id)
	if err != nil {
		logrus.Errorf("ERROR: don't get users %s", err)
		return post, err
	}
	logrus.Info("DONE:get user data from psql")

	return post, err
}
