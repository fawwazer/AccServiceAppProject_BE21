package users

import (
	"fmt"

	"gorm.io/gorm"
)

type Users struct {
	HP           uint `gorm:"primaryKey"`
	Nama         string
	Password     string
	Alamat       string
	Userbalances []UserBalance
}

type UserBalance struct {
	gorm.Model
	Balance   int64  //nilai saldo sekarang
	UsersID   uint   // foreign key users
	Transaksi string //keterangan top up / transfer
	Nilai     int64  //histori berapa uang keluar atau masuk
}

func (ub *UserBalance) GetBalance(connection *gorm.DB, userid uint) (UserBalance, error) {
	// var result uint
	var user UserBalance
	err := connection.Where("users_id = ?", userid).Last(&user).Error
	if err != nil {
		return UserBalance{}, err
	}
	return user, nil
}

func (ub *UserBalance) TopUp(connection *gorm.DB, hp uint, amount int64) (bool, error) {
	// var user UserBalance
	ub.Balance += amount // balance terbaru ditambah amount akan tersimpan di balance
	// query := connection.Table("user_balances").Where(Users{"hp = ?", hp}).Update("balance", ub.Balance)
	// query := connection.Model(&Users{HP: hp}).Updates(UserBalance{Balance: updateBalance})
	// if connection.Table("user_balances").Where("users_id = ? AND transaksi = ?", hp, "topup").First(&result) {

	// }
	// if err := connection.Table("user_balances").Where("users_id = ? AND transaksi = ?", hp, "topup").First(&user).Error; err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		user := UserBalance{Balance: updateBalance, UsersID: hp, Transaksi: "topup", Nilai: amount}
	// 		query := connection.Select("Balance", "UsersID", "Transaksi", "Nilai").Create(&user)
	// 		if err := query.Error; err != nil {
	// 			return false, err
	// 		}
	// 		return query.RowsAffected > 0, nil
	// 	} else {
	// 		user := UserBalance{Balance: updateBalance, Nilai: amount}
	// 		query := connection.Model(&user).Updates(UserBalance{Balance: updateBalance, Nilai: amount})
	// 		if err := query.Error; err != nil {
	// 			return false, err
	// 		}
	// 		return query.RowsAffected > 0, nil
	// 	}
	// }

	user := UserBalance{Balance: ub.Balance, UsersID: hp, Transaksi: "topup", Nilai: amount}
	query := connection.Select("Balance", "UsersID", "Transaksi", "Nilai").Create(&user)
	if err := query.Error; err != nil {
		return false, err
	}

	return query.RowsAffected > 0, nil
}

func (u *Users) Transfer(connection *gorm.DB, user1 *Users, user2 *Users, balance int64) {

	var user1Balance, user2Balance int64

	// Retrieve balance for user1
	connection.Model(user1).Association("Userbalances").Find(&user1.Userbalances)
	for _, ub := range user1.Userbalances {
		user1Balance += ub.Balance
	}

	// Retrieve balance for user2
	connection.Model(user2).Association("Userbalances").Find(&user2.Userbalances)
	for _, ub := range user2.Userbalances {
		user2Balance += ub.Balance
	}

	// Subtract user2's balance from user1's balance
	resultbalance := user1Balance - user2Balance

	// Now resultbalance contains the difference between user1's and user2's balances
	fmt.Printf("Resulting balance after transfer: %d\n", resultbalance)
}

func Login(connection *gorm.DB, hp uint, password string) (Users, error) {
	var result Users
	err := connection.Where("hp = ? AND password = ?", hp, password).First(&result).Error
	if err != nil {
		return Users{}, err
	}

	return result, nil
}

func Register(connection *gorm.DB, newUser Users) (bool, error) {
	query := connection.Create(&newUser)
	if err := query.Error; err != nil {
		return false, err
	}

	return query.RowsAffected > 0, nil
}

func (u *Users) GantiPassword(connection *gorm.DB, newPassword string) (bool, error) {
	query := connection.Table("users").Where("hp = ?", u.HP).Update("password", newPassword)
	if err := query.Error; err != nil {
		return false, err
	}

	return query.RowsAffected > 0, nil
}

func (u *Users) DeleteAcc(connection *gorm.DB, hp uint) (bool, error) {
	query := connection.Table("users").Where("hp = ?", hp).Delete(hp)
	if err := query.Error; err != nil {
		return false, err
	}

	return query.RowsAffected > 0, nil
}

func SeeAnotherAcc(connection *gorm.DB, hp string) (Users, error) {
	var result Users
	err := connection.Where("hp = ?", hp).First(&result).Error
	if err != nil {
		return Users{}, err
	}

	return result, nil
}
