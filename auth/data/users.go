package data

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username           string             `json:"username" bson:"email" validate:"required,email"`
	Password           string             `json:"password" bson:"password" validate:"required,min=6,max=20"`
	jwt.StandardClaims `json:"-"`
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
