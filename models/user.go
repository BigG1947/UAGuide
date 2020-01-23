package models

import (
	"crypto/sha256"
	"database/sql"
)

type User struct {
	Id       int
	FIO      string
	Email    string
	Phone    string
	Password []byte
	Photo    string
}

func (u *User) Add(db *sql.DB) error {
	u.Password = generatePasswordHash(u.Password)
	res, err := db.Exec("INSERT INTO users(fio, email, phone, password) VALUES ($1, $2, $3, $4)", u.FIO, u.Email, u.Phone, u.Password)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	u.Id = int(id)
	return nil
}

func (u *User) Update(db *sql.DB) error {
	if _, err := db.Exec("UPDATE users SET fio = $1, email = $2, phone = $3, password = $4, photo = $5 WHERE id = $6", u.FIO, u.Email, u.Phone, u.Password, u.Photo, u.Id); err != nil {
		return err
	}
	return nil
}

func (u *User) Get(db *sql.DB) error {
	if err := db.QueryRow("SELECT id, fio, email, phone, password, photo FROM users WHERE id = $1", u.Id).Scan(&u.Id, &u.FIO, &u.Email, &u.Phone, &u.Password, &u.Photo); err != nil {
		return err
	}
	return nil
}

func (u *User) Login(db *sql.DB) error {
	u.Password = generatePasswordHash(u.Password)
	if err := db.QueryRow("SELECT id, fio, email, phone, password, photo FROM users WHERE email = $1 AND password = $2", u.Email, u.Password).Scan(&u.Id, &u.FIO, &u.Email, &u.Phone, &u.Password, &u.Photo); err != nil {
		return err
	}
	return nil
}

func generatePasswordHash(password []byte) []byte {
	passwordHash := sha256.Sum256(password)
	return passwordHash[:]
}

func CheckUserPhoneExist(phone string, db *sql.DB) bool {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM users WHERE phone = ?;", phone).Scan(&count)
	if count > 0 {
		return false
	}
	return true
}

func CheckUserEmailExist(email string, db *sql.DB) bool {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?;", email).Scan(&count)
	if count > 0 {
		return false
	}
	return true
}
