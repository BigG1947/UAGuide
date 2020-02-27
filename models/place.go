package models

import (
	"database/sql"
)

type Place struct {
	Id          int
	Name        string
	City        City
	Photo       string
	PrevPhoto   string
	Description string
	Status      int
	Private     bool
	User        User
	FB          []PlaceFeedBack
	PP          []PlacePhoto
	Date        string
}

type PlaceFeedBack struct {
	Id      int
	User    User
	Comment string
	Mark    int
	Date    string
}

type PlacePhoto struct {
	Id      int
	Src     string
	Private bool
	Date    string
}

type PlaceList []Place

func (p *Place) Add(db *sql.DB) error {
	if _, err := db.Exec("INSERT INTO places(name, description, cities_id, user_id, photo, photo_prev, status, private, date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", p.Name, p.Description, p.City.Id, p.User.Id, p.Photo, p.PrevPhoto, p.Status, p.Private, p.Date); err != nil {
		return err
	}
	return nil
}

func (pl *PlaceList) GetAllCount(db *sql.DB) (int, error) {
	return 0, nil
}
func (pl *PlaceList) GetItemsForPage(db *sql.DB, page int, limit int) error {
	return nil
}
func (pl *PlaceList) GetItemsForPageByParameters(db *sql.DB, page int, limit int, parameters map[string]string) error {
	return nil
}
func (pl *PlaceList) GetCountForUser(db *sql.DB, id int) (int, error) {
	return 0, nil
}
func (pl *PlaceList) GetItemsForPageForUser(db *sql.DB, page int, id int) error {
	return nil
}
func (pl *PlaceList) GetItemsForPageForUserByParameters(db *sql.DB, page int, limit int, parameters map[string]string, id int) error {
	return nil
}
