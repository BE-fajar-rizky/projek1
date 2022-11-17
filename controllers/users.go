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
	statm := db.QueryRow("SELECT Id_user,phone,kata_sandi FROM user WHERE phone = ? AND kata_sandi = ?", user.Phone, user.Kata_sandi)

	var row entity.User
	errs := statm.Scan(&row.Id, &row.Phone, &row.Kata_sandi)
	//  bcrypt.GenerateFromPassword(, bcrypt.DefaultCost))

	if errs != nil {
		log.Fatal("Maaf No Telfon atau Password salah ")
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
func UpdateUser(db *sql.DB, update entity.User, Phone string) (sql.Result, error) {
	// res := db.QueryRow("SELECT Id_user,Nama_user,phone,alamat,foto_profil from user where id_user=?", id)
	// var barisUser entity.User
	var query = "UPDATE user set Nama_user = ?, email = ?, phone = ?,alamat = ? ,kata_sandi = ?  where phone = ? "
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("erorr prepare update", errPrepare.Error())
	}
	result, errExec := statement.Exec(update.Nama, update.Email, update.Phone, update.Alamat, update.Kata_sandi, Phone)

	if errExec != nil {
		log.Fatal("erorr Exec update", errExec.Error())
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
func Top_up(db *sql.DB, total entity.Top_up, Id int) (entity.User, error) {
	usr := db.QueryRow("SELECT Id_user, Nama_user,phone from user where phone=?", total.Phone)

	var rowUser entity.User
	errscan := usr.Scan(&rowUser.Id, &rowUser.Nama, &rowUser.Phone)
	fmt.Println(rowUser.Id)
	var query = "INSERT INTO topup(users_Id,Jumlah_Topup) VALUES (?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("erorr prepare insert", errPrepare.Error())

	}

	result, errExec := statement.Exec(Id, total.Jumlah_TopUP)
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

	if errscan != nil {
		if errscan == sql.ErrNoRows {
			log.Fatal("error scan", errscan.Error())
		}
	}
	return rowUser, nil

}
func TfUser(db *sql.DB, total entity.Transfers, id int) (entity.User, error) {
	usr := db.QueryRow("SELECT Id_user, Nama_user,phone from user where phone=?", total.Phone)
	// usr2 := db.QueryRow("SELECT Id_user, Nama_user,phone from user where id=?", id)
	// var barisUser entity.User
	// errscan2:= usr2.Scan(&barisUser.Id,&barisUser.Nama)
	var rowUser entity.User
	errscan := usr.Scan(&rowUser.Id, &rowUser.Nama, &rowUser.Phone)
	fmt.Println(rowUser.Id)
	var query = "INSERT INTO transfers(pengirim_id,Jumlah_TF,penerima_id) VALUES (?,?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("erorr prepare insert", errPrepare.Error())

	}

	result, errExec := statement.Exec(id, total.Jumlah_TF, rowUser.Id)
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

	if errscan != nil {
		if errscan == sql.ErrNoRows {
			log.Fatal("error scan", errscan.Error())
		}
	}
	return rowUser, nil

}
func DeleteAkun(db *sql.DB, phone string) (sql.Result, error) {
	var query = "DELETE FROM user where phone = ?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("prepare error", errPrepare.Error())
	}
	result, ErrExec := statement.Exec(phone)
	if ErrExec != nil {
		log.Fatal("error exec", ErrExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("Delete berhasil")
		} else {
			fmt.Println("Delete gagal")
		}
	}

	return result, nil
}
