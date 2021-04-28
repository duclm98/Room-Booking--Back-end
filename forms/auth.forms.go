package forms

import (
	dto "echo-demo/DTOs"
)

type Auth struct {
	AccessToken  string `param:"accessToken" query:"accessToken" json:"accessToken"`
	RefreshToken string `param:"refreshToken" query:"refreshToken" json:"refreshToken"`
	User         dto.User   `param:"user" query:"user" json:"user"`
}
