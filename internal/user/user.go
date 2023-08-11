package user

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type User struct {
	ID        int
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) encryptPassword(password string) {
	hash := sha256.New()
	hash.Write([]byte(password))
	hashed := hash.Sum(nil)
	u.Password = fmt.Sprintf("%x", hashed)
}
