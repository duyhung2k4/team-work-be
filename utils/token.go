package utils

import (
	"context"
	"net/http"
	"strings"
	"team-work-be/config"
)

type FieldToken string

const (
	USER_ID  FieldToken = "user_id"
	ROLE     FieldToken = "role"
	USERNAME FieldToken = "username"
)

func GetFieldToken(field FieldToken, r *http.Request) (interface{}, error) {
	tokenCookie := r.Header.Get("Authorization")

	token, errToken := config.GetJWT().Decode(strings.Split(tokenCookie, " ")[1])

	if errToken != nil {
		return nil, errToken
	}

	claims, err := token.AsMap(context.Background())

	if err != nil {
		return nil, err
	}

	return claims[string(field)], nil
}
