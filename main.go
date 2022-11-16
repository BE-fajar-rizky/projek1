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
			var row entity.User
			//buat mekanisme login

			fmt.Println("masukkan nama user")
			fmt.Scanln(&row.Phone)
			fmt.Println("masukkan Kata Sandi anda")
			fmt.Scanln(&row.Kata_sandi)
			dataLogin, err := controllers.LoginUser(db, row)
			if err != nil {
				log.Fatal("error gagal", err.Error())
			}
			fmt.Println("Login berhasil", dataLogin)
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
			updateUser := entity.User{}
			fmt.Println("masukkan id user yang akan diupdate :")
			fmt.Scanln(&updateUser.Id)
			fmt.Println("masukkan nama user")
			fmt.Scanln(&updateUser.Nama)
			fmt.Println("masukkan Email user")
			fmt.Scanln(&updateUser.Email)
			fmt.Println("masukkan Phone user")
			fmt.Scanln(&updateUser.Phone)
			fmt.Println("masukkan alamat user")
			fmt.Scanln(&updateUser.Alamat)
			fmt.Println("masukkan Foto profil user")
			fmt.Scanln(&updateUser.Foto_profil)
			fmt.Println("masukkan kata sandi user")
			fmt.Scanln(&updateUser.Kata_sandi)

			dataUser, errUpdate := controllers.UpdateUser(db, updateUser)
			if errUpdate != nil {
				log.Fatal("error gagal", errUpdate.Error())
			}
			fmt.Println("Update berhasil", dataUser)
		}

	case 5:
		{
			fmt.Println("baca data by id")
		}

	}

}
