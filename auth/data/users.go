package data

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (u *User) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func (u *User) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) Validate() error {
	validate := validator.New()

	return validate.Struct(u)
}
