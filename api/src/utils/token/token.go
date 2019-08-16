package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nsysu/teacher-education/src/persistence/redis"
	"github.com/nsysu/teacher-education/src/utils/config"
	"github.com/nsysu/teacher-education/src/utils/hash"
	uuid "github.com/satori/go.uuid"
)

type claims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}

// AccessToken set claims from parameters and config file and returns access token built
func AccessToken(params map[string]string) (string, error) {
	jti := uuid.NewV4().String()
	now := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		Account: params["account"],
		StandardClaims: jwt.StandardClaims{
			Issuer:    config.Get("jwt.issuer").(string),
			IssuedAt:  now,
			NotBefore: now,
			ExpiresAt: now + int64(config.Get("jwt.access_token_exp").(int)),
			Id:        jti,
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

// Parse validates token gotten from request and returns claims if it's legal
func Parse(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.Get("jwt.secret").(string)), nil
	})

	if err != nil {
		return nil, err
	}

	claims, err := validate(token)

	if err != nil {
		return nil, err
	}
	return claims, err
}

func validate(token *jwt.Token) (map[string]interface{}, error) {
	conn := redis.Redis()
	defer conn.Close()

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok {
		if claims["iss"] != config.Get("jwt.issuer") {
			return nil, errors.New("token issuer mismatch")
		}

		user := redis.UserDao.Get(conn, claims["account"].(string))
		if user.JTI != claims["jti"].(string) {
			return nil, errors.New("token id mismatch")
		}
	} else {
		return nil, errors.New("get token claims failed")
	}

	return map[string]interface{}{
		"account": claims["account"],
	}, nil
}
