package advertisements

import (
	"bytes"
	"crypto/sha256"
	"errors"
)

type User struct {
	name        string
	password    []byte
	permissions int32
}

var ErrPassword = errPassword()

func errPassword() error {
	return errors.New("пароль должен состоять минимум из 3 символов")
}

func NewUser(name, password string, needSalt bool, perm int32) (User, error) {
	u := User{}
	err := u.SetName(name)
	if err != nil {
		return User{}, err
	}
	err = u.SetPassword(password, needSalt)
	if err != nil {
		return User{}, err
	}
	u.SetPermissions(perm)
	return u, nil
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) error {
	if len(name) < 2 {
		return errClientName
	}
	u.name = name
	return nil
}

func (u *User) Password() []byte {
	return u.password
}

func (u *User) SetPassword(password string, needSalt bool) error {
	if len(password) == 0 {
		return nil
	}
	if needSalt {
		u.password = u.saltPassword(password)
		return nil
	}
	u.password = []byte(password)
	return nil
}

func (u *User) saltPassword(pass string) []byte {
	var b bytes.Buffer
	nameSumm := sha256.Sum256([]byte(u.name))
	passSumm := sha256.Sum256([]byte(pass))
	b.Grow(len(nameSumm) + len(passSumm))
	for i := 0; i < len(nameSumm); i++ {
		if i%2 == 0 {
			b.Write([]byte{nameSumm[i], passSumm[len(passSumm)-1-i]})
			continue
		}
		b.Write([]byte{passSumm[i], nameSumm[len(nameSumm)-1-i]})
	}

	p := sha256.Sum256(b.Bytes())
	return p[:]
}

func (u *User) Permissions() int32 {
	return u.permissions
}

func (u *User) SetPermissions(p int32) {
	u.permissions = p
}
