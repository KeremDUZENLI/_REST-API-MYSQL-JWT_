package mapper

import (
	"jwt-project/database/model"
	"jwt-project/dto"
)

func MapperSignUp(d *dto.DtoSignUp) model.Person {
	return model.Person{
		ID:           d.ID,
		Password:     d.Password,
		Token:        d.Token,
		RefreshToken: d.RefreshToken,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
		UserId:       d.UserId,

		FirstName: d.FirstName,
		LastName:  d.LastName,
		Email:     d.Email,
		UserType:  d.UserType,
	}
}

func MapperLogin(d *dto.DtoLogIn) model.Person {
	return model.Person{
		Password: d.Password,
		Email:    d.Email,

		ID:           d.ID,
		Token:        d.Token,
		RefreshToken: d.RefreshToken,
		UpdatedAt:    d.UpdatedAt,
		UserId:       d.UserId,
	}
}

func MapperGetUser(d *dto.GetUser) model.Person {
	return model.Person{
		UserId: d.UserId,

		ID:           d.ID,
		FirstName:    d.FirstName,
		LastName:     d.LastName,
		Password:     d.Password,
		Email:        d.Email,
		UserType:     d.UserType,
		Token:        d.Token,
		RefreshToken: d.RefreshToken,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
	}
}
