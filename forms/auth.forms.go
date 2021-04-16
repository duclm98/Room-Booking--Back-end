package forms

type Auth struct {
	AccessToken  string `param:"accessToken" query:"accessToken" json:"accessToken"`
	RefreshToken string `param:"refreshToken" query:"refreshToken" json:"refreshToken"`
	User         User
}
