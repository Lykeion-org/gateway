package models


type LoginRequest struct {
	Username 		string		`json:"username"`
	Password 		string		`json:"password"`
}

type TokenResponse struct {
	AccessToken		string		`json:"accessToken"`
	RefreshToken	string		`json:"refreshToken"`
}

type CheckAccessLevelRequest struct {
	RequiredLevel	int32		`json:"requiredLevel"`
}


