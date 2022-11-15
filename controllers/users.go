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
func loginUser(db *sql.DB, id int, Phone string, kata_sandi string) (entity.User, error) {
	res := db.QueryRow("SELECT Id_user,Nama_user,phone,alamat,foto_profil, kata_sandi from user where id_user=?", id, Phone, kata_sandi)
	var barisUser entity.User

	errscan := res.Scan(&barisUser.Phone, &barisUser.kata_sandi)
	statement, errPrepare := db.Prepare(errscan)

	if errPrepare != nil {
		log.Fatal("erorr prepare insert", errPrepare.Error())
	}
	result, errExec := statement.Exec(loginUser.Phone, loginUser.kata_sandi)
	if errExec != nil {
		log.Fatal("erorr Exec insert", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("LOGIN berhasil")
		} else {
			fmt.Println("LOGIN gagal")
		}
	}
	return result, nil
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
func updateUser(db *sql.DB, newUser entity.User) (sql.Result, error) {
	res := db.QueryRow("SELECT Id_user,Nama_user,phone,alamat,foto_profil from user where id_user=?", id)
	var barisUser entity.User
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
