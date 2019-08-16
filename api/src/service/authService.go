package service

import (
	"github.com/nsysu/teacher-education/src/error"
	"github.com/nsysu/teacher-education/src/persistence/gorm"
	"github.com/nsysu/teacher-education/src/utils/config"
	"github.com/nsysu/teacher-education/src/utils/hash"
	"github.com/nsysu/teacher-education/src/utils/logger"
	"github.com/nsysu/teacher-education/src/utils/token"
)

// Login user login
func Login(account, password string) (result map[string]interface{}, e *error.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = error.UnexpectedError()
		}
	}()

	user := gorm.UserDao.GetByAccount(tx, account)

	if user == nil {
		return nil, error.LoginError()
	}

	if ok := hash.Verify(password, user.Password); !ok {
		return nil, error.LoginError()
	}

	accessToken, err := token.AccessToken(map[string]string{
		"account": user.Account,
	})
	if err != nil {
		panic(err)
	}

	refreshToken, err := token.RefreshToken(user.Account)
	if err != nil {
		panic(err)
	}

	result = map[string]interface{}{
		"Token":        accessToken,
		"RefreshToken": refreshToken,
		"Expire":       config.Get("jwt.access_token_exp").(int),
	}
	return result, nil
}
