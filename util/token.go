package util

import (
	"errors"
	"fmt"
	"log"
	"os"
	"social_network/constant/env"
	"social_network/constant/noti"
	"strings"
	"time"
	"unicode"

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

//func ExtractDataFromToken(tokenString string, logger *log.Logger) (string, string, time.Time, error) {
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, errors.New("Unexpected signing method")
//		}
//
//		return []byte(os.Getenv(env.SECRET_KEY)), nil
//	})
//
//	if err != nil {
//		logger.Println("Error at ExtractDataFromToken - ", err)
//		return "", "", time.Time{}, errors.New(noti.GenericsErrorWarnMsg)
//	}
//
//	var userId string = ""
//	var role string = ""
//	var exp time.Time
//
//	claims, ok := token.Claims.(jwt.MapClaims)
//
//	if !ok || !token.Valid {
//		return "", "", time.Time{}, errors.New(noti.GenericsErrorWarnMsg)
//	}
//
//	if rawRole, ok := claims["role"].(string); rawRole == "" || !ok {
//		return "", "", time.Time{}, errors.New(noti.GenericsErrorWarnMsg)
//	} else {
//		role = rawRole
//	}
//
//	if id, ok := claims["userId"].(string); id == "" || !ok {
//		return "", "", time.Time{}, errors.New(noti.GenericsErrorWarnMsg)
//	} else {
//		userId = id
//	}
//
//	if expPeriod, ok := claims["exp"].(time.Time); !ok {
//		return "", "", time.Time{}, errors.New(noti.GenericsErrorWarnMsg)
//	} else {
//		exp = expPeriod
//	}
//
//	return userId, role, exp, nil
//}

func ExtractDataFromToken(tokenString string, logger *log.Logger) (string, string, time.Time, error) {
	// Check for empty token
	if tokenString == "" {
		logger.Println("Error at ExtractDataFromToken - empty token")
		return "", "", time.Time{}, errors.New(noti.GenericsErrorWarnMsg)
	}

	// Remove "Bearer " prefix if present
	if len(tokenString) > 7 && strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimSpace(tokenString[7:])
	}

	// Trim any whitespace and non-printable characters
	tokenString = strings.TrimSpace(tokenString)

	// Check if token contains invalid characters
	for i, c := range tokenString {
		if !unicode.IsPrint(c) || c == ' ' {
			logger.Printf("Error at ExtractDataFromToken - invalid character at position %d: %q", i, c)
			return "", "", time.Time{}, errors.New(noti.GenericsErrorWarnMsg)
		}
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secretKey := os.Getenv(env.SECRET_KEY)
		if secretKey == "" {
			logger.Println("Error at ExtractDataFromToken - secret key not found in environment")
			return nil, errors.New("missing secret key")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		logger.Printf("Error at ExtractDataFromToken - %v", err)
		return "", "", time.Time{}, errors.New(noti.GenericsErrorWarnMsg)
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		logger.Println("Error at ExtractDataFromToken - invalid claims or token")
		return "", "", time.Time{}, errors.New(noti.GenericsErrorWarnMsg)
	}

	// Extract user_id
	userId, ok := claims["user_id"].(string)
	if !ok || userId == "" {
		logger.Println("Error at ExtractDataFromToken - missing or invalid user_id claim")
		return "", "", time.Time{}, errors.New(noti.GenericsErrorWarnMsg)
	}

	// Extract role
	role, ok := claims["role"].(string)
	if !ok || role == "" {
		logger.Println("Error at ExtractDataFromToken - missing or invalid role claim")
		return "", "", time.Time{}, errors.New(noti.GenericsErrorWarnMsg)
	}

	// Extract expiration
	expFloat, ok := claims["expire"].(float64)
	if !ok {
		logger.Println("Error at ExtractDataFromToken - missing or invalid expiration claim")
		return "", "", time.Time{}, errors.New(noti.GenericsErrorWarnMsg)
	}

	// Convert Unix timestamp to time.Time
	exp := time.Unix(int64(expFloat), 0)

	return userId, role, exp, nil
}
