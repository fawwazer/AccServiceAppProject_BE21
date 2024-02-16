package main

import (
	"AccServiceProject_BE21/config"
	"AccServiceProject_BE21/users"
	"fmt"
)

func main() {
	database := config.InitMysql()
	config.Migrate(database)
	connection := database
	var input int

	for input != 99 {
		fmt.Println("Pilih menu")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
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
				fmt.Println("Terjadi kesalahan (tidak bisa mendaftarkan pengguna)", err.Error())
			}
			if success {
				fmt.Println("Selamat anda telah terdaftar")
			}
		case 2:
			var hp, password string
			fmt.Println("Masukkan HP")
			fmt.Scanln(&hp)
			fmt.Println("Masukkan Password")
			fmt.Scanln(&password)
			loggedIn, err := users.Login(database, hp, password)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			fmt.Println("Selamat Datang,", loggedIn.Nama)

			// Exit case 2 block after successful login
		caseLoop:
			for {
				var choice int
				fmt.Println("1. My Account")
				fmt.Println("2. Update Account")
				fmt.Println("3. Delete Account")
				fmt.Println("4. Top Up")
				fmt.Println("5. Transfer")
				fmt.Println("6. History Top Up")
				fmt.Println("7. History Transfer")
				fmt.Println("8. See Another User")
				fmt.Println("0. Log Out")
				fmt.Print("Masukkan pilihan:")
				fmt.Scanln(&choice)
				switch choice {
				case 1:
					// My Account logic
					balance := users.SumBalance(loggedIn.Userbalances)
					fmt.Println("Nama:", loggedIn.Nama)
					fmt.Println("Nomor HP:", loggedIn.HP)
					fmt.Println("Alamat:", loggedIn.Alamat)
					fmt.Println("Balance:", balance)

					var isRunning bool = true
					for isRunning {
						var inputLogin int
						fmt.Println("Silahkan kembali ke menu dengan mengetik angka 1")
						fmt.Print("Masukkan angka:")
						fmt.Scanln(&inputLogin)
						if inputLogin == 1 {
							isRunning = false
						}
					}
				case 2:
					// Update Account logic
					// Implement the logic here
					var newh users.Users
					fmt.Print("Nomor HP: ", loggedIn.HP)
					newh.HP = loggedIn.HP
					fmt.Print("Masukkan Password Baru: ")
					fmt.Scanln(&newh.Password)
					fmt.Print("Masukkan nama baru: ")
					fmt.Scanln(&newh.Nama)
					fmt.Print("Masukkan alamat baru: ")
					fmt.Scanln(&newh.Alamat)

					success, err := newh.UpdateAcc(database, newh.HP, newh.Password, newh.Nama, newh.Alamat)
					if err != nil {
						fmt.Println("Terjadi kesalahan (tidak bisa mengupdate akun):", err.Error())
					}
					if success {
						fmt.Println("Selamat, Akun Telah Terupdate")
					}
				case 3:
					// Delete Account logic
					// Implement the logic here
					var deleteAccount users.Users
					var pilih int8
					deleteAccount.HP = loggedIn.HP
					fmt.Println("1. Yes")
					fmt.Println("2. No")
					fmt.Scanln(&pilih)
					if pilih == 1 {
						success, err := deleteAccount.DeleteAcc(database, deleteAccount.HP)

						if err != nil {
							fmt.Println("Terjadi kesalahan(tidak bisa mendaftarkan pengguna)", err.Error())
						}
						if success {
							fmt.Println("Account Telah Dihapus!")
						}
					} else {
						var isRunning bool = true
						for isRunning {
							var inputLogin int
							fmt.Println("Silahkan kembali ke menu dengan mengetik angka 1")
							fmt.Print("Masukkan angka:")
							fmt.Scanln(&inputLogin)
							if inputLogin == 1 {
								isRunning = false
							}
						}
					}

				case 4:
					// Top Up logic
					// Implement the logic here
					var hp uint
					var nominal int64
					var ub users.UserBalance
					fmt.Print("Nomor Hp: ", loggedIn.HP)
					hp = loggedIn.HP
					fmt.Print("Masukkan Nominal: ")
					fmt.Scanln(&nominal)
					success, err := ub.TopUp(database, hp, nominal)
					if err != nil {
						fmt.Println("terjadi kesalahan(tidak bisa top up)", err.Error())
					} else if success {
						fmt.Println("Top up Berhasil")
					}
				case 5:
					// Transfer logic
					// Implement the logic here
					var transfers uint
					var transfers2 uint
					var transferb int64
					fmt.Print("hp user saya :", loggedIn.HP)
					transfers = loggedIn.HP
					fmt.Print("masukan hp user penerima :")
					fmt.Scanln(&transfers2)
					fmt.Print("masukan jumlah transfer :")
					fmt.Scanln(&transferb)

					user1, err := users.GetUserByHP(database, transfers)
					if err != nil {
						fmt.Println("Error:", err)
						return
					}
					user2, err := users.GetUserByHP(database, transfers2)
					if err != nil {
						fmt.Println("Error:", err)
						return
					}

					success, err := users.Transfer(connection, user1, user2, transferb)
					if err != nil {
						fmt.Println("Error:", err)
						return
					}

					if success {
						fmt.Println("Transfer successful!")
					} else {
						fmt.Println("Transfer failed.")
					}
				case 6:
					// History Top Up logic
					// Implement the logic here
					var isRunning bool = true
					for isRunning {
						var hp uint
						var History users.Users
						fmt.Println("No HP: ", loggedIn.HP)
						hp = loggedIn.HP
						History, balance, err := users.SeeAnotherAcc(database, hp)
						if err == nil {
							balance = users.SumBalance(History.Userbalances)
							fmt.Println("Berikut Data User Tersebut: ")
							fmt.Println("Nama: ", History.Nama)
							fmt.Println("Nomor HP: ", History.HP)
							fmt.Println("Alamat: ", History.Alamat)
							fmt.Println("Balance:", balance)
							fmt.Println()
							balance, err := users.Historytopup(database, hp)
							if err == nil {
								var inputLogin int
								fmt.Println("Berikut Data History User Tersebut: ")
								fmt.Println("Balance:", balance)
								fmt.Println()
								fmt.Println("Silahkan kembali ke menu dengan mengetik angka 1")
								fmt.Print("Masukkan angka:")
								fmt.Scanln(&inputLogin)
								if inputLogin == 1 {
									isRunning = false
								}
							}
						}
					}
				case 7:
					// History Transfer logic
					// Implement the logic here
					var isRunning bool = true
					for isRunning {
						var hp uint
						var History users.Users
						fmt.Println("Untuk Melihat User Lain Mohon Input No HP")
						hp = loggedIn.HP
						History, balance, err := users.SeeAnotherAcc(database, hp)
						if err == nil {
							balance = users.SumBalance(History.Userbalances)
							fmt.Println("Berikut Data User Tersebut: ")
							fmt.Println("Nama: ", History.Nama)
							fmt.Println("Nomor HP: ", History.HP)
							fmt.Println("Alamat: ", History.Alamat)
							fmt.Println("Balance:", balance)
							fmt.Println()
							balance, err := users.Historytransfer(database, hp)
							if err == nil {
								var inputLogin int
								fmt.Println("Berikut Data History User Tersebut: ")
								fmt.Println("Balance:", balance)
								fmt.Println()
								fmt.Println("Silahkan kembali ke menu dengan mengetik angka 1")
								fmt.Print("Masukkan angka:")
								fmt.Scanln(&inputLogin)
								if inputLogin == 1 {
									isRunning = false
								}
							}
						}
					}
				case 8:
					// See Another User logic
					// Implement the logic here
					var isRunning bool = true
					for isRunning {
						var hp uint
						var seeAcc users.Users
						fmt.Println("Untuk Melihat User Lain Mohon Input No HP")
						fmt.Println("Masukkan No HP: ")
						fmt.Scanln(&hp)
						seeAcc, balance, err := users.SeeAnotherAcc(database, hp)
						if err == nil {
							var inputLogin int
							fmt.Println("Berikut Data User Tersebut: ")
							fmt.Println("Nama: ", seeAcc.Nama)
							fmt.Println("Nomor HP: ", seeAcc.HP)
							fmt.Println("Alamat: ", seeAcc.Alamat)
							fmt.Println("Balance:", balance)
							fmt.Println("Silahkan kembali ke menu dengan mengetik angka 1")
							fmt.Print("Masukkan angka:")
							fmt.Scanln(&inputLogin)
							if inputLogin == 1 {
								isRunning = false
							}
						}
					}
				case 0:
					fmt.Println("Log Out...")
					break caseLoop // Exit caseLoop and continue executing subsequent cases
				default:
					fmt.Println("Pilihan tidak valid.")
				}
			}
		case 0:
			fmt.Println("Exiting...")
			input = 99 // Exit caseLoop and continue executing subsequent cases
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
