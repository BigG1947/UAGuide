package models

import "database/sql"

type City struct {
	Id     int
	Name   string
	Region Region
}

type Region struct {
	Id   int
	Name string
}

type CitiesList []City

func (cl *CitiesList) Get(db *sql.DB) error {
	rows, err := db.Query("SELECT cities.id, cities.name, regions.id, regions.name FROM cities, regions WHERE regions.id = cities.id_regions")
	if err != nil {
		return err
	}
	for rows.Next() {
		var c City
		if err := rows.Scan(&c.Id, &c.Name, &c.Region.Id, &c.Region.Name); err != nil {
			return err
		}
		*cl = append(*cl, c)
	}
	return nil
}
