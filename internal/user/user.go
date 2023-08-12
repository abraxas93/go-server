package user

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) EncryptPassword() {
	hash := sha256.New()
	hash.Write([]byte(u.Password))
	hashed := hash.Sum(nil)
	u.Password = fmt.Sprintf("%x", hashed)
}
