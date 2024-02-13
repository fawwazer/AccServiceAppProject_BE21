package main

import (
	"AccServiceProject_BE21/config"
	"AccServiceProject_BE21/users"

	"fmt"
)

func main() {
	database := config.InitMysql()
	config.Migrate(database)
	var input int
	for input != 99 {
		fmt.Println("Pilih menu")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. My Account")
		fmt.Println("4. Update Account")
		fmt.Println("5. Delete Account")
		fmt.Println("6. Top Up")
		fmt.Println("7. Transfer")
		fmt.Println("8. History Top Up")
		fmt.Println("9. History Transfer")
		fmt.Println("10. See Another User")
		fmt.Println("0. Exit")
		fmt.Print("Masukkan pilihan:")
		fmt.Scanln(&input)
		switch input {
		case 1:
			var newUser users.Users
			fmt.Print("Masukkan Nama: ")
			fmt.Scanln(&newUser.Nama)
			fmt.Print("Masukkan Nomor HP: ")
			fmt.Scanln(&newUser.HP)
			fmt.Print("Masukkan Password: ")
			fmt.Scanln(&newUser.Password)
			fmt.Print("Masukkan Alamat: ")
			fmt.Scanln(&newUser.Alamat)
			success, err := users.Register(database, newUser)
			if err != nil {
				fmt.Println("terjadi kesalahan(tidak bisa mendaftarkan pengguna)", err.Error())
			}
			if success {
				fmt.Println("selamat anda telah terdaftar")
			}
		case 2:
			var isRunning bool = true
			for isRunning {
				var hp string
				var password string
				var loggedIn users.Users
				fmt.Println("Masukkan HP")
				fmt.Scanln(&hp)
				fmt.Println("Masukkan Password")
				fmt.Scanln(&password)
				loggedIn, err := users.Login(database, hp, password)
				if err == nil {
					var inputLogin int
					fmt.Println("Selamat Datang,", loggedIn.Nama)
					fmt.Println("Pilih Menu Kamu")
					fmt.Println("1. Logout")
					fmt.Print("Masukkan pilihan:")
					fmt.Scanln(&inputLogin)
					if inputLogin == 1 {
						isRunning = false
					}
				}
			}
		case 3:
			var isRunning bool = true
			for isRunning {
				var hp string
				var seeAcc users.Users
				fmt.Println("Untuk Melihat User Lain Mohon Input No HP")
				fmt.Println("Masukkan No HP: ")
				fmt.Scanln(&hp)
				seeAcc, err := users.SeeAnotherAcc(database, hp)
				if err == nil {
					var inputLogin int
					fmt.Println("Berikut Data User Tersebut: ")
					fmt.Println("Nama: ", seeAcc.Nama)
					fmt.Println("Nomor HP: ", seeAcc.HP)
					fmt.Println("Alamat: ", seeAcc.Alamat)
					fmt.Println("Silahkan kembali ke menu dengan mengetik angka 1")
					fmt.Print("Masukkan angka:")
					fmt.Scanln(&inputLogin)
					if inputLogin == 1 {
						isRunning = false
					}
				}
			}
		case 4:
			var newPassword users.Users
			fmt.Print("Masukkan Nomor HP: ")
			fmt.Scanln(&newPassword.HP)
			fmt.Print("Masukkan Password Baru: ")
			fmt.Scanln(&newPassword.Password)
			success, err := newPassword.GantiPassword(database, newPassword.Password)
			if err != nil {
				fmt.Println("Terjadi kesalahan(tidak bisa mendaftarkan pengguna)", err.Error())
			}
			if success {
				fmt.Println("Selamat Account Telah Terupdate")
			}
		case 5:
			var deleteAccount users.Users
			fmt.Print("Masukkan Nomor HP: ")
			fmt.Scanln(&deleteAccount.HP)
			success, err := deleteAccount.DeleteAcc(database, deleteAccount.HP)
			if err != nil {
				fmt.Println("Terjadi kesalahan(tidak bisa mendaftarkan pengguna)", err.Error())
			}
			if success {
				fmt.Println("Account Telah Dihapus!")
			}
		case 6:
			fmt.Println("HEHe")
		case 7:
			fmt.Println("HEHe")
		case 8:
			fmt.Println("HEHe")
		case 9:
			fmt.Println("HEHe")
		case 10:
			var isRunning bool = true
			for isRunning {
				var hp string
				var seeAcc users.Users
				fmt.Println("Untuk Melihat User Lain Mohon Input No HP")
				fmt.Println("Masukkan No HP: ")
				fmt.Scanln(&hp)
				seeAcc, err := users.SeeAnotherAcc(database, hp)
				if err == nil {
					var inputLogin int
					fmt.Println("Berikut Data User Tersebut: ")
					fmt.Println("Nama: ", seeAcc.Nama)
					fmt.Println("Nomor HP: ", seeAcc.HP)
					fmt.Println("Alamat: ", seeAcc.Alamat)
					fmt.Println("Silahkan kembali ke menu dengan mengetik angka 1")
					fmt.Print("Masukkan angka:")
					fmt.Scanln(&inputLogin)
					if inputLogin == 1 {
						isRunning = false
					}
				}
			}
		case 0:

			fmt.Println("HEHe")
		}
	}
	// fmt.Println("Exited! Thank you")

}
