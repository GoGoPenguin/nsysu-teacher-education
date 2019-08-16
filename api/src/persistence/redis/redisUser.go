package redis

import (
	"fmt"
	"reflect"

	"github.com/nsysu/teacher-education/src/utils/config"
)

// User model
type User struct {
	Account      string `redis:"-"`
	JTI          string `redis:"JTI"`
	RefreshToken string `redis:"REFRESH_TOKEN"`
}

type userDao struct{}

// UserDao user data acces object
var UserDao userDao

func (*userDao) Get(conn *Conn, account string) *User {
	user := User{Account: account}

	v := reflect.ValueOf(&user).Elem()
	t := reflect.TypeOf(user)

	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)

		if tag, ok := f.Tag.Lookup("redis"); ok && tag != "-" {
			key := fmt.Sprintf("%s-%s", user.Account, tag)

			value, err := conn.get(key)
			if err != nil {
				panic(err)
			}
			v.Field(i).SetString(value)
		}
	}
	return &user
}

func (*userDao) Store(conn *Conn, user *User) {
	v := reflect.ValueOf(*user)
	t := reflect.TypeOf(*user)
	set := fmt.Sprintf("%s-SET", user.Account)

	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)

		if tag, ok := f.Tag.Lookup("redis"); ok && tag != "-" {
			key := fmt.Sprintf("%s-%s", user.Account, tag)
			ttl := config.Get("jwt.access_token_exp").(int)
			if tag == "REFRESH_TOKEN" {
				ttl = config.Get("jwt.refresh_token_exp").(int)
			}

			if err := conn.sAdd(set, key); err != nil {
				panic(err)
			}
			if err := conn.setEx(key, ttl, v.FieldByName(f.Name).String()); err != nil {
				panic(err)
			}
		}
	}

	if err := conn.expire(set, config.Get("jwt.refresh_token_exp").(int)); err != nil {
		panic(err)
	}
}

func (*userDao) Delete(conn *Conn, account string) {
	key := fmt.Sprintf("%s-SET", account)

	keys, err := conn.sMembers(key)
	if err != nil {
		panic(err)
	}

	keys = append(keys, key)
	if err := conn.del(keys); err != nil {
		panic(err)
	}
}
