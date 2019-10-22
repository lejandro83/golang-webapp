package models

import (
	"errors"
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidLogin = errors.New("invalid login")
)

func AuthenticateUser(usr, pwd string) error {
	hash, err := client.Get("user:" + usr).Bytes()
	if err == redis.Nil {
		return ErrUserNotFound
	} else if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(pwd))
	if err != nil {
		return ErrInvalidLogin
	}
	return nil
}

func RegisterUser(usr, pwd string) error {
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	if err != nil {
		return err
	}
	return client.Set("user:"+usr, hash, 0).Err()
}
