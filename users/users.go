package users

import "gorm.io/gorm"

type Users struct {
	HP       string
	Nama     string
	Password string
	Alamat   string
}

type UserBalance struct {
	gorm.Model
	Balance 	int64
	UserID 		uint
}

type History struct {
	gorm.Model
	Transaksi		string
	Nilai			int64
	Detail			string
	UserBalanceID 	uint
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

func (u *Users) DeleteAcc(connection *gorm.DB, hp string) (bool, error) {
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
