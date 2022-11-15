package main

import (
	config "fajars/rowsql/config"
	"fajars/rowsql/controllers"
	entity "fajars/rowsql/entity"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.KonekDB()

	defer db.Close() // menutup koneksi

	//buat mekanisme menu
	fmt.Println("MENU:\n1. register\n2. Login\n3. read\n4. update\n5. delete ")
	fmt.Println("Masukkan pilihan anda:")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		{
			// dataUser, errGetAll := controllers.Userdata(db) // memanggil fungsi di controllers
			// if errGetAll != nil {                           // cek apakah ada error dari getAllUser
			// 	fmt.Println("error get all users")
			// }
			// // fmt.Println("data all", dataUser)
			// for _, value := range dataUser { // membaca/menampilkan  seluruh data user yang telah ditampung di variable slice
			// 	fmt.Printf("id: %s, nama: %s, email: %s, phone: %s\n", value.Id, value.Nama, value.Gender, value.Status)
			// }
			newUser := entity.User{}

			fmt.Println("masukkan nama user")
			fmt.Scanln(&newUser.Nama)
			fmt.Println("masukkan email anda")
			fmt.Scanln(&newUser.Email)
			fmt.Println("masukkan No Hp anda")
			fmt.Scanln(&newUser.Phone)
			fmt.Println("masukkan alamat anda")
			fmt.Scanln(&newUser.Alamat)
			fmt.Println("masukkan foto anda")
			fmt.Scanln(&newUser.Foto_profil)
			fmt.Println("masukkan kata sandi anda")
			fmt.Scanln(&newUser.Kata_sandi)

			controllers.Register(db, newUser)
		}
	case 2:
		{
			var phone int
			fmt.Println("silahkan masukkan username anda:")
			fmt.Scan(&phone)
			var kata_sandi int
			fmt.Println("silahkan masukkan username anda:")
			fmt.Scan(&kata_sandi)

			dataUser, errGetId := controllers.loginUser(db, phone, kata_sandi)
			loginUser := entity.User{}
			//buat mekanisme login
			fmt.Println("LOGIN:\n1. username\n2. Password")
			fmt.Println("Masukkan akun anda:")
			fmt.Println("1.masukkan username anda")
			fmt.Scanln(&loginUser.Phone)
			fmt.Println("masukkan kata sandi anda")
			fmt.Scanln(&loginUser.kata_sandi)
			var query = "LOGIN Users set Phone = ?, kata sandi = ? "
			statement, errPrepare := db.Prepare(query)
			if errPrepare != nil {
				log.Fatal("error prepare", errPrepare.Error())
			}
			result, errExec := statement.Exec(loginUser.Phone, loginUser.kata_sandi)
			if errExec != nil {
				log.Fatal("error exec", errExec.Error())
			} else {
				row, _ := result.RowsAffected()
				if row > 0 {
					fmt.Println("LOGIN berhasil")
				} else {
					fmt.Println("LOGIN gagal")
				}
			}
			fmt.Println("LOGIN")
		}
	case 3:
		{
			var id int
			fmt.Println("silahkan masukkan id")
			fmt.Scan(&id)

			dataUser, errGetId := controllers.Readsdata(db, id)

			if errGetId != nil {
				log.Fatal("id tidak ada", errGetId.Error())
			}

			fmt.Printf("id : %d\n Nama User : %s \n Phone : %s\n Alamat : %s\n foto_profil :", dataUser.Id, dataUser.Nama, dataUser.Phone, dataUser.Alamat, dataUser.Foto_profil)
		}

	case 4:
		{
			updateUser := entity.user{}
			fmt.Println("masukkan id user yang akan diupdate :")
			fmt.Scanln(&updateUser.Id_user)
			fmt.Println("masukkan nama user")
			fmt.Scanln(&updateUser.Nama_user)
			fmt.Println("masukkan Email user")
			fmt.Scanln(&updateUser.Email)
			fmt.Println("masukkan Phone user")
			fmt.Scanln(&updateUser.Phone)
			fmt.Println("masukkan alamat user")
			fmt.Scanln(&updateUser.alamat)
			fmt.Println("masukkan Foto profil user")
			fmt.Scanln(&updateUser.foto_profil)
			fmt.Println("masukkan kata sandi user")
			fmt.Scanln(&updateUser.kata_sandi)

			var query = "UPDATE Users set name = ?, Email = ?, Phone = ?, Alamat = ?, Foto Profil = ?, kata sandi = ? "
			statement, errPrepare := db.Prepare(query)
			if errPrepare != nil {
				log.Fatal("error prepare", errPrepare.Error())
			}
			result, errExec := statement.Exec(updateUser.Id_user, updateUser.Nama_user, updateUser.Email, updateUser.Phone, updateUser.alamat, updateUser.foto_profil, updateUser.kata_sandi)
			if errExec != nil {
				log.Fatal("error exec", errExec.Error())
			} else {
				row, _ := result.RowsAffected()
				if row > 0 {
					fmt.Println("UPDATE berhasil")
				} else {
					fmt.Println("UPDATE gagal")
				}
			}
			fmt.Println("update")
		}

	case 5:
		{
			fmt.Println("baca data by id")
		}

	}

}
