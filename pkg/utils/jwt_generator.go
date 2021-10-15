package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)


type Tokens struct {
	Access  string
	Refresh string
}


func GenerateNewTokens(id string, credentials []string) (*Tokens, error) {

	accessToken, err := generateNewAccessToken(id, credentials)
	if err != nil {

		return nil, err
	}

	refreshToken, err := generateNewRefreshToken()
	if err != nil {
		return nil, err
	}

	return &Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func generateNewAccessToken(id string, credentials []string) (string, error) {

	secret := os.Getenv("JWT_SECRET_KEY")


	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	claims := jwt.MapClaims{}

	claims["id"] = id
	claims["expires"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	claims["product:create"] = false
	claims["product:update"] = false
	claims["product:delete"] = false
	claims["category:create"] = false
	claims["category:update"] = false
	claims["category:delete"] = false
	for _, credential := range credentials {
		claims[credential] = true
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}

func generateNewRefreshToken() (string, error) {

	hash := sha256.New()


	refresh := os.Getenv("JWT_REFRESH_KEY") + time.Now().String()

	//https://pkg.go.dev/io#Writer.Write
	_, err := hash.Write([]byte(refresh))
	if err != nil {
		return "", err
	}


	hoursCount, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT"))


	expireTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix())


	t := hex.EncodeToString(hash.Sum(nil)) + "." + expireTime

	return t, nil
}

func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}
