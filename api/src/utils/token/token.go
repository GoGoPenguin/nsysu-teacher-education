package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/persistence/redis"
	"github.com/nsysu/teacher-education/src/utils/config"
	"github.com/nsysu/teacher-education/src/utils/hash"
)

type claims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}

// AccessToken set claims from parameters and config file and returns access token built
func AccessToken(params map[string]string) (string, error) {
	now := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		Account: params["account"],
		StandardClaims: jwt.StandardClaims{
			Issuer:    config.Get("jwt.issuer").(string),
			IssuedAt:  now,
			NotBefore: now,
			ExpiresAt: now + int64(config.Get("jwt.access_token_exp").(int)),
			Id:        params["jti"],
		},
	})

	// issues : https://github.com/dgrijalva/jwt-go/issues/65
	// `SignedString` is receiving `[]byte`, but it declares it accepts `interface{}`
	result, err := token.SignedString([]byte(config.Get("jwt.secret").(string)))
	if err != nil {
		return "", err
	}

	return result, nil
}

// RefreshToken combine user account and timestamp to generate refresh token
func RefreshToken(account string) (string, error) {
	result := hash.New(account + time.Now().String())
	return result, nil
}

// Validate token gotten from request and returns claims if it's legal
func Validate(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.Get("jwt.secret").(string)), nil
	})

	if err != nil {
		return err
	}

	if err := checkClaims(token); err != nil {
		return err
	}
	return nil
}

// Claims get claims from token in header
func Claims(ctx iris.Context) map[string]interface{} {
	tokenString, _ := jwtmiddleware.FromAuthHeader(ctx)
	token, _ := jwt.Parse(tokenString, nil)
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims
}

func checkClaims(token *jwt.Token) error {
	conn := redis.Redis()
	defer conn.Close()

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok {
		if claims["iss"] != config.Get("jwt.issuer") {
			return errors.New("token issuer mismatch")
		}

		user := redis.UserDao.Get(conn, claims["account"].(string))
		if user.JTI != claims["jti"].(string) {
			return errors.New("token id mismatch")
		}
	} else {
		return errors.New("get token claims failed")
	}

	return nil
}
