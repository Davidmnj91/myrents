package domain

import (
	"errors"
	"github.com/Davidmnj91/myrents/pkg/types"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	UserUUID  types.UUID
	Username  string
	Password  string
	Name      string
	Surname   string
	IDNumber  string
	Email     string
	Phone     string
	BirthDate types.Date
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (u *User) Create() error {
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(password)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	return nil
}

func (u *User) Delete() {
	u.DeletedAt = time.Now()
}

func (u *User) Empty() bool {
	return len(u.IDNumber) == 0 && len(u.Email) == 0 && len(u.Username) == 0 && len(u.Phone) == 0
}

func (u *User) Verify(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}
