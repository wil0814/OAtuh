package jwt

import (
	"errors"
	"go-OAuth/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Key cookie key
const Key = "token"

// Claims Token的結構，裡面放你要的資訊
type Claims struct {
	Username string `json:"user_name"`
	UserID   string `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken 產生Token
func GenerateToken(userID, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(config.Val.JWTTokenLife) * time.Second) // Token有效時間

	claims := Claims{
		username,
		userID,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "william",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(config.Val.JWTSecret))

	return token, err
}

// ParseToken 驗證Token對不對，如果對就回傳user info
func ParseToken(token string) (userID, userName string, err error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Val.JWTSecret), nil
	})

	if err != nil {
		return "", "", err
	}

	claims, ok := tokenClaims.Claims.(*Claims)
	if !ok || !tokenClaims.Valid {
		return "", "", errors.New("tokenClaims invalid")
	}

	return claims.UserID, claims.Username, nil
}
