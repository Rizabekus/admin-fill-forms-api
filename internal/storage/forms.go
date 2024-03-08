package storage

import (
	"database/sql"
	"fmt"
)

type FormsDB struct {
	DB *sql.DB
}

func CreateFormsStorage(db *sql.DB) *FormsDB {
	return &FormsDB{DB: db}
}
func (fdb *FormsDB) CreateSession(uuid string) error {
	stmt, err := fdb.DB.Prepare("INSERT INTO admin_sessions(session) VALUES($1)")
	if err != nil {
		fmt.Println("qweqwe")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(uuid)
	if err != nil {

		return err
	}
	return nil
}
func (fdb *FormsDB) CheckSession(uuid string) (bool, error) {
	query := "SELECT COUNT(*) FROM admin_sessions WHERE session = $1"

	var count int
	err := fdb.DB.QueryRow(query, uuid).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {

			return false, nil
		}

		return false, err
	}

	return count > 0, nil
}
