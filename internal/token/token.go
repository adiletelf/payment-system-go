package token

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(admin_id uint) (string, string, error) {
	accessTokenLifespan, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_MINUTE_LIFESPAN"))
	if err != nil {
		return "", "", err
	}
	refreshTokenLifespan, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", "", err
	}

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["admin_id"] = admin_id
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(accessTokenLifespan)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := at.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return "", "", err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["admin_id"] = admin_id
	rtClaims["exp"] = time.Now().Add(time.Hour * time.Duration(refreshTokenLifespan)).Unix()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshToken, err := rt.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
func TokenValid(r *http.Request) error {
	tokenString := ExtractToken(r, "token")
	_, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(r *http.Request, tokenName string) string {
	token := r.URL.Query().Get(tokenName)
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractTokenID(r *http.Request) (uint, error) {
	tokenString := ExtractToken(r, "token")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["admin_id"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}

func RefreshAccessRefreshTokens(r *http.Request) (string, string, error) {
	refreshToken := ExtractToken(r, "refreshToken")
	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return "", "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		admin_id, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["admin_id"]), 10, 64)
		if err != nil {
			return "", "", err
		}
		at, rt, err := GenerateToken(uint(admin_id))
		if err != nil {
			return "", "", err
		}
		return at, rt, nil
	}
	return "", "", nil
}
