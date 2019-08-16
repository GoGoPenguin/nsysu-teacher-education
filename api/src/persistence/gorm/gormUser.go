package gorm

import (
	"github.com/jinzhu/gorm"
)

const (
	table = "user"
)

// User model
type User struct {
	gorm.Model
	Name     string `gorm:"column:name;"`
	Account  string `gorm:"column:account; unique_index"`
	Password string `gorm:"column:password;"`
	Role     string `gorm:"column:role; default:'student'"`
}

type userDao struct{}

// UserDao user data acces object
var UserDao userDao

// New a record
func (*userDao) New(tx *gorm.DB, user *User) {
	err := tx.Table(table).
		Create(user).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (*userDao) GetByID(tx *gorm.DB, id uint) *User {
	result := User{}
	err := tx.Table(table).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return &result
}

// GetByAccount get a record by id
func (*userDao) GetByAccount(tx *gorm.DB, acount string) *User {
	result := User{}
	err := tx.Table(table).
		Where("account = ?", acount).
		Where("deleted_at IS NULL").
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return &result
}

// Query custom query
func (*userDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]User {
	var result []User
	err := tx.Table(table).
		Scopes(funcs...).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return &result
}
