package jwt_utils

import "LoanTrackerApi/internal/entity"

func RefreshToken(tokenString string) (entity.Token, error) {

	claims, err := ValidateToken(tokenString)

	if err != nil {
		return entity.Token{}, err
	}

	access, err := CreateToken(*claims)
	if err != nil {
		return entity.Token{}, err
	}
	refresh, err := CreateToken(*claims)

	if err != nil {
		return entity.Token{}, err
	}

	return entity.Token{AccessToken: access, RefreshToken: refresh}, nil

}
