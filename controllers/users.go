package controllers

import (
	"database/sql"
	entity "fajars/rowsql/entity"
	"fmt"
	"log"
)

func Register(db *sql.DB, newUser entity.User) (sql.Result, error) {

	var query = "INSERT INTO user(nama_user,email,phone,alamat,foto_profil,kata_sandi) VALUES (?,?,?,?,?,?)"
	statement, errPrepare := db.Prepare(query)

	if errPrepare != nil {
		log.Fatal("erorr prepare insert", errPrepare.Error())

	}
	result, errExec := statement.Exec(newUser.Nama, newUser.Email, newUser.Phone, newUser.Alamat, newUser.Foto_profil, newUser.Kata_sandi)
	if errExec != nil {
		log.Fatal("erorr Exec insert", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("berhasil")
		} else {
			fmt.Println("gagal")
		}
	}
	return result, nil
}
func LoginUser(db *sql.DB, user entity.User) (entity.User, error) {
	statm, err := db.Query("SELECT Id, Phone, kata_sandi FROM users WHERE phone = ? AND kata_sandi = ?") //PREPARE MENYIAPKAN KODE YG AKAN DI EKSEKUSI DI SQL
	if err != nil {
		log.Fatal("error select ", err.Error())
	}
	var row entity.User
	for statm.Next() {
		errs := statm.Scan(&row.Phone, &row.Kata_sandi)
		if errs != nil {
			log.Fatal("error scan ", errs.Error())
		}
	}
	return row, nil
}

func Readsdata(db *sql.DB, id int) (entity.User, error) {
	res := db.QueryRow("SELECT Id_user,Nama_user,phone,alamat,foto_profil from user where id_user=?", id)
	var barisUser entity.User

	errscan := res.Scan(&barisUser.Id, &barisUser.Nama, &barisUser.Phone, &barisUser.Alamat, &barisUser.Foto_profil)

	if errscan != nil {
		if errscan == sql.ErrNoRows {
			log.Fatal("error scan", errscan.Error())
		}
	}
	return barisUser, nil
}
func UpdateUser(db *sql.DB, newUser entity.User) (sql.Result, error) {
	// res := db.QueryRow("SELECT Id_user,Nama_user,phone,alamat,foto_profil from user where id_user=?", id)
	// var barisUser entity.User
	var query = "INSERT INTO user(nama_user,email,phone,alamat,foto_profil,kata_sandi) VALUES (?,?,?,?,?,?)"
	statement, errPrepare := db.Prepare(query)

	if errPrepare != nil {
		log.Fatal("erorr prepare insert", errPrepare.Error())

	}
	result, errExec := statement.Exec(newUser.Nama, newUser.Email, newUser.Phone, newUser.Alamat, newUser.Foto_profil, newUser.Kata_sandi)
	if errExec != nil {
		log.Fatal("erorr Exec insert", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("berhasil")
		} else {
			fmt.Println("gagal")
		}
	}
	return result, nil
}
