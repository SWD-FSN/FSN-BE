package util

import (
	"errors"
	"fmt"
	"log"
	"os"
	"social_network/constant/env"
	"social_network/constant/noti"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokens(email, userId, role string, logger *log.Logger) (string, string, error) {
	var bytes = []byte(os.Getenv(env.SECRET_KEY))
	var errMsg string = "Error while generating tokens - "

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"role":    role,
		"expire":  time.Now().Add(AccessDuration).Unix(),
	}).SignedString(bytes)
	if err != nil {
		logger.Print(errMsg + fmt.Sprint(err))
		return "", "", errors.New(noti.InternalErr)
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"role":    role,
		"expire":  time.Now().Add(RefreshDuration).Unix(),
	}).SignedString(bytes)
	if err != nil {
		logger.Print(errMsg + fmt.Sprint(err))
		return "", "", errors.New(noti.InternalErr)
	}

	return accessToken, refreshToken, nil
}

func GenerateActionToken(email, userId, role string, logger *log.Logger) (string, error) {
	var bytes = []byte(os.Getenv(env.SECRET_KEY))
	var errMsg string = "Error while generating action token - "

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"role":    role,
		"expire":  time.Now().Add(NormalActionDuration).Unix(),
	}).SignedString(bytes)
	if err != nil {
		logger.Print(errMsg + fmt.Sprint(err))
		return "", errors.New(noti.InternalErr)
	}

	return token, nil
}

func ExtractDataFromToken(tokenString string, logger *log.Logger) (string, string, time.Time, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(os.Getenv(env.SECRET_KEY)), nil
	})

	if err != nil {
		logger.Println("Error at ExtractDataFromToken - ", err)
		return "", "", time.Time{}, errors.New(noti.GenericsErrorWarnMsg)
	}

	var userId string = ""
	var role string = ""
	var exp time.Time

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return "", "", time.Time{}, errors.New(noti.GenericsErrorWarnMsg)
	}

	if rawRole, ok := claims["role"].(string); rawRole == "" || !ok {
		return "", "", time.Time{}, errors.New(noti.GenericsErrorWarnMsg)
	} else {
		role = rawRole
	}

	if id, ok := claims["userId"].(string); id == "" || !ok {
		return "", "", time.Time{}, errors.New(noti.GenericsErrorWarnMsg)
	} else {
		userId = id
	}

	if expPeriod, ok := claims["exp"].(time.Time); !ok {
		return "", "", time.Time{}, errors.New(noti.GenericsErrorWarnMsg)
	} else {
		exp = expPeriod
	}

	return userId, role, exp, nil
}
