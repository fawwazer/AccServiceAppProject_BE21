package users

import (
	"errors"

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
func GetUserByHP(connection *gorm.DB, hp string) (*Users, error) {
	var user Users
	if err := connection.Where("hp = ?", hp).Preload("Userbalances").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ub *UserBalance) TopUp(connection *gorm.DB, hp uint, amount int64) (bool, error) {
	ub.Balance += amount // Increment the balance by the top-up amount

	// Create a new balance record
	transaction := UserBalance{
		Balance:   ub.Balance,
		UsersID:   hp,
		Transaksi: "topup",
		Nilai:     amount,
	}

	// Create the balance record in the database
	if err := connection.Create(&transaction).Error; err != nil {
		return false, err
	}

	return true, nil
}

func Transfer(connection *gorm.DB, user1 *Users, user2 *Users, amount int64) (bool, error) {
	if user1 == nil || user2 == nil {
		return false, errors.New("invalid user")
	}
	if user1.HP == user2.HP {
		return false, errors.New("cannot transfer to the same user")
	}

	// Retrieve the current balance records of user1 and user2
	var user1Balance UserBalance
	if err := connection.Where("users_id = ?", user1.HP).Last(&user1Balance).Error; err != nil {
		return false, err
	}

	var user2Balance UserBalance
	if err := connection.Where("users_id = ?", user2.HP).Last(&user2Balance).Error; err != nil {
		return false, err
	}

	// Check if user1 has sufficient balance for the transfer
	if user1Balance.Balance < amount {
		return false, errors.New("insufficient balance")
	}

	// Perform the transfer by updating the balances
	user1Balance.Balance -= amount
	user2Balance.Balance += amount

	// Create transaction records for user1 and user2
	transaction1 := UserBalance{
		Balance:   user1Balance.Balance,
		UsersID:   user1.HP,
		Transaksi: "transfer",
		Nilai:     -amount,
	}
	transaction2 := UserBalance{
		Balance:   user2Balance.Balance,
		UsersID:   user2.HP,
		Transaksi: "transfer",
		Nilai:     amount,
	}

	// Begin a database transaction
	tx := connection.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Update user1 balance
	if err := tx.Create(&transaction1).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	// Update user2 balance
	if err := tx.Create(&transaction2).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return false, err
	}

	return true, nil
}

func Login(connection *gorm.DB, hp string, password string) (Users, error) {
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

func SeeAnotherAcc(connection *gorm.DB, hp string) (Users, int64, error) {
	var user Users

	// Fetch the user's data including their balance records
	err := connection.Where("hp = ?", hp).Preload("Userbalances").First(&user).Error
	if err != nil {
		return Users{}, 0, err
	}

	// If the user has balance records, return the current balance
	if len(user.Userbalances) > 0 {
		return user, user.Userbalances[0].Balance, nil
	}

	// Return the user's data and 0 balance if no balance records are found
	return user, 0, nil
}
