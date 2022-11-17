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
	fmt.Println("MENU:\n1. register\n2. Login  ")
	fmt.Println("Masukkan pilihan anda:")
	var pilihan int
	fmt.Scanln(&pilihan)
	var data entity.User
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
			var err error
			fmt.Println("masukkan phone anda")
			fmt.Scanln(&row.Phone)
			fmt.Println("masukkan Kata Sandi anda")
			fmt.Scanln(&row.Kata_sandi)
			data, err = controllers.LoginUser(db, row)
			if err != nil {
				log.Fatal("error gagal", err.Error())
			}
			fmt.Println("Login berhasil", data)
		}
		fmt.Println("halo menu utama 2")
		runis := true
		for runis {
			fmt.Print("SUB MENU:\n1. Read Account \n2. Update Account \n3. Delete Account \n4. Top-Up \n5. Transfer \n6. History Top-Up \n7. History Transfer \n8. Melihat Profil user lain \n9. Keluar\n")
			fmt.Println("Masukkan Pilihan Anda ")
			var choice2 int
			fmt.Scanln(&choice2)
			switch choice2 {
			case 1:
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
			case 2:
				{
					updateUser := entity.User{}
					fmt.Println("masukkan id user yang akan diupdate :")
					fmt.Scanln(&updateUser.Id)
					fmt.Println("masukkan nama user")
					fmt.Scanln(&updateUser.Nama)
					fmt.Println("masukkan Email user")
					fmt.Scanln(&updateUser.Email)
					fmt.Println("masukkan phone")
					fmt.Scanln(&updateUser.Phone)
					fmt.Println("masukkan alamat user")
					fmt.Scanln(&updateUser.Alamat)
					fmt.Println("masukkan kata sandi user")
					fmt.Scanln(&updateUser.Kata_sandi)

					dataUser, errUpdate := controllers.UpdateUser(db, updateUser)
					if errUpdate != nil {
						log.Fatal("error gagal", errUpdate.Error())
					}
					fmt.Println("Update berhasil", dataUser)
				}
			case 3:
				{

					// transfer := entity.Transfers{}
					// fmt.Println("masukkan jumlah transfer :")
					// fmt.Scanln(&transfer.Jumlah_TF)
					// fmt.Println("masukkan nama hanphone yang akan di kirim")
					// fmt.Scanln(&transfer.Phone)

					// controllers.TfUser(db, transfer, data.Id)
				}
			case 4:
				{
					TopUp := entity.Top_up{}
					fmt.Println("masukkan jumlah topup :")
					fmt.Scanln(&TopUp.Jumlah_TopUP)
					fmt.Println("masukkan nomor hanphone yang akan di topup")
					fmt.Scanln(&TopUp.Phone)

					controllers.Top_up(db, TopUp, data.Id)
				}
			case 5:
				{
					transfer := entity.Transfers{}
					fmt.Println("masukkan jumlah transfer :")
					fmt.Scanln(&transfer.Jumlah_TF)
					fmt.Println("masukkan nama hanphone yang akan di kirim")
					fmt.Scanln(&transfer.Phone)

					controllers.TfUser(db, transfer, data.Id)
				}
				// case 3:
				// 	{
				// 		fmt.Println("HALO SUB MENU 3")
				// 	}
				// case 4:
				// 	{
				// 		fmt.Println("HALO SUB MENU 4")
				// 	}
				// case 5:
				// 	{
				// 		transfer := entity.Transfers{}
				// 		fmt.Println("masukkan jumlah transfer :")
				// 		fmt.Scanln(&transfer.Jumlah_TF)
				// 		fmt.Println("masukkan nama hanphone yang akan di kirim")
				// 		fmt.Scanln(&transfer.Phone)

				// 		controllers.TfUser(db, transfer, data.Id)
				// 	}
			}
		}
	}
}
