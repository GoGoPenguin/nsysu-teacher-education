package service

import (
	"github.com/nsysu/teacher-education/src/error"
	"github.com/nsysu/teacher-education/src/persistence/gorm"
	"github.com/nsysu/teacher-education/src/persistence/redis"
	"github.com/nsysu/teacher-education/src/utils/config"
	"github.com/nsysu/teacher-education/src/utils/hash"
	"github.com/nsysu/teacher-education/src/utils/logger"
	"github.com/nsysu/teacher-education/src/utils/token"
	uuid "github.com/satori/go.uuid"
)

// Login user login
func Login(account, password, role string) (result interface{}, e *error.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = error.UnexpectedError()
		}
	}()

	var (
		userAccount  string
		userPassword string
	)

	if role == gorm.AdminDao.Role {
		admin := gorm.AdminDao.GetByAccount(tx, account)
		if admin == nil {
			return nil, error.LoginError()
		}

		userAccount = admin.Account
		userPassword = admin.Password
	} else {
		student := gorm.StudentDao.GetByName(tx, account)
		if student == nil {
			return nil, error.LoginError()
		}

		userAccount = student.Account
		userPassword = student.Password
	}

	if ok := hash.Verify(password, userPassword); !ok {
		return nil, error.LoginError()
	}

	jti := uuid.NewV4().String()
	accessToken, err := token.AccessToken(map[string]string{
		"account": userAccount,
		"jti":     jti,
	})
	if err != nil {
		panic(err)
	}

	refreshToken, err := token.RefreshToken(userAccount)
	if err != nil {
		panic(err)
	}

	conn := redis.Redis()
	defer conn.Close()

	redisUser := &redis.User{
		Account:      userAccount,
		JTI:          jti,
		RefreshToken: refreshToken,
	}
	redis.UserDao.Store(conn, redisUser)

	result = map[string]interface{}{
		"Token":        accessToken,
		"RefreshToken": refreshToken,
		"Expire":       config.Get("jwt.access_token_exp").(int),
	}
	return result, nil
}

// Logout user logout
func Logout(account string) (result string, e *error.Error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = error.UnexpectedError()
		}
	}()

	conn := redis.Redis()
	defer conn.Close()

	redis.UserDao.Delete(conn, account)

	return "success", nil
}

// RenewToken get new access token
func RenewToken(account, refreshToken string) (result map[string]interface{}, e *error.Error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = error.UnexpectedError()
		}
	}()

	conn := redis.Redis()
	defer conn.Close()

	redisUser := redis.UserDao.Get(conn, account)

	// validate refresh token
	if redisUser.RefreshToken != refreshToken {
		return nil, error.ValidateError("Invalid refresh token")
	}

	// if user renew token before token expired, it means token is stolen.
	if redisUser.JTI != "" {
		redis.UserDao.Delete(conn, account)
		return nil, error.RevokeTokenError()
	}

	// create a new access token
	jti := uuid.NewV4().String()
	accessToken, err := token.AccessToken(map[string]string{
		"account": account,
		"jti":     jti,
	})
	if err != nil {
		panic(err)
	}

	// update redis data
	redisUser = &redis.User{
		Account:      account,
		JTI:          jti,
		RefreshToken: refreshToken,
	}
	redis.UserDao.Store(conn, redisUser)

	result = map[string]interface{}{
		"Token":  accessToken,
		"Expire": config.Get("jwt.access_token_exp").(int),
	}
	return result, nil
}
