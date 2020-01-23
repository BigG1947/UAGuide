package models

import (
	"database/sql"
)

type Place struct {
	Id          int
	Name        string
	Photo       string
	Description string
	FB          []PlaceFeedBack
}

type PlaceFeedBack struct {
	Id      int
	User    User
	Comment string
	Mark    int
	Date    string
}

type PlaceList []Place

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
