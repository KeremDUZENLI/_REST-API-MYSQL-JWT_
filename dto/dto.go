package dto

import (
	"context"
	"time"

	"jwt-project/common/constants"
	"jwt-project/database"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type DtoSignUp struct {
	ID           primitive.ObjectID `bson:"_id"`
	Password     string             `json:"password" validate:"required,min=6"`
	Token        string             `json:"token"`
	RefreshToken string             `json:"refreshtoken"`
	CreatedAt    time.Time          `json:"createdat"`
	UpdatedAt    time.Time          `json:"updatedat"`
	UserId       string             `json:"userid"`

	FirstName string `json:"firstname" validate:"required,min=2,max=100"`
	LastName  string `json:"lastname" validate:"required,min=2,max=100"`
	Email     string `json:"email" validate:"email,required"`
	UserType  string `json:"usertype" validate:"required,eq=ADMIN|eq=USER"`
}

func (d DtoSignUp) IsNotExist(ctx context.Context) bool { return !Amount(ctx, d) }

func (d DtoSignUp) IsObeyRules() bool { return Validator(d) == nil }

func Amount(ctx context.Context, person DtoSignUp) bool {
	number, err := database.Collection(database.Connect(), constants.TABLE).CountDocuments(ctx, bson.M{"email": person.Email})
	if err != nil {
		return false
	}

	return number > 0
}

func Validator(d DtoSignUp) error {
	return validator.New().Struct(d)
}

// ----------------------------------------------------------------

type DtoLogIn struct {
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"email,required"`

	ID           primitive.ObjectID `bson:"_id"`
	Token        string             `json:"token"`
	RefreshToken string             `json:"refreshtoken"`
	UpdatedAt    time.Time          `json:"updatedat"`
	UserId       string             `json:"userid"`
}

func (d DtoLogIn) IsValidEmail(email string) bool { return email == d.Email }

func (d DtoLogIn) IsValidPassword(password string) bool { return Verify(password, d.Password) }

func Verify(password string, providedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}

func Find(ctx context.Context, d DtoLogIn) *DtoLogIn {
	var foundPerson DtoLogIn
	if err := database.Collection(database.Connect(), constants.TABLE).FindOne(ctx, bson.M{"email": d.Email}).Decode(&foundPerson); err != nil {
		return &d
	}
	return &foundPerson
}

// ----------------------------------------------------------------

type GetUser struct {
	UserId string `json:"userid"`

	ID           primitive.ObjectID `bson:"_id"`
	FirstName    string             `json:"firstname" validate:"required,min=2,max=100"`
	LastName     string             `json:"lastname" validate:"required,min=2,max=100"`
	Password     string             `json:"password" validate:"required,min=6"`
	Email        string             `json:"email" validate:"email,required"`
	UserType     string             `json:"usertype" validate:"required,eq=ADMIN|eq=USER"`
	Token        string             `json:"token"`
	RefreshToken string             `json:"refreshtoken"`
	CreatedAt    time.Time          `json:"createdat"`
	UpdatedAt    time.Time          `json:"updatedat"`
}
