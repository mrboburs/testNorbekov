package repository

import (
	"fmt"
	"github.com/mrboburs/Norbekov/model"
	"github.com/mrboburs/Norbekov/util/logrus"

	"github.com/jmoiron/sqlx"
)

type ContactPostDB struct {
	db *sqlx.DB
}

func NewContactPostDB(db *sqlx.DB) *ContactPostDB {
	return &ContactPostDB{db: db}
}

func (repo *ContactPostDB) CreateContactPost(post model.Contact, logrus *logrus.Logger) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (first_name ,last_name  , phone_number , type_service,text) VALUES ($1, $2, $3,$4,$5)  RETURNING id", contact)

	row := repo.db.QueryRow(query, post.FirstName, post.LastName, post.PhoneNumber, post.TypeService, post.Text)

	if err := row.Scan(&id); err != nil {
		logrus.Infof("ERROR:PSQL Insert error %s", err.Error())
		return 0, err
	}
	logrus.Info("DONE: INSERTED Data PSQL")
	return id, nil
}
func (repo *ContactPostDB) GetAllContact(logrus *logrus.Logger) (contact []model.ContactFull, err error) {
	rowsRs, err := repo.db.Query("SELECT id,first_name,last_name,phone_number,type_service,text,created_at  FROM contact")

	if err != nil {
		logrus.Infof("ERROR: not selecting data from sql %s", err.Error())
		// http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return contact, err
	}

	contactArray := []model.ContactFull{}
	defer rowsRs.Close()

	for rowsRs.Next() {
		snb := model.ContactFull{}
		err = rowsRs.Scan(&snb.ID, &snb.FirstName, &snb.LastName, &snb.PhoneNumber, &snb.TypeService, &snb.Text, &snb.Created_At)
		if err != nil {
			logrus.Infof("ERROR: not scanning data from sql %s", err.Error())
			// log.Println(err)
			// http.Error(w, http.StatusText(500), 500)
			return contact, err
		}
		contactArray = append(contactArray, snb)
	}

	if err = rowsRs.Err(); err != nil {

		return contactArray, err
	}
	return contactArray, nil
}
