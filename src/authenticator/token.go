package authenticator

import (
	"errors"
	"io/ioutil"
	"time"

	cfg "kumu-exam/config"

	"github.com/golang-jwt/jwt/v4"
)

const (
	prvKey = "storage/keys/private_key.pem"
	pubKey = "storage/keys/public_key.pem"
)

func GenerateToken() (string, error) {
	bPrvKey, err := ioutil.ReadFile(prvKey)
	if err != nil {
		return "", errors.New(cfg.InternalError)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(bPrvKey)
	if err != nil {
		return "", errors.New(cfg.InternalError)
	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["vld"] = true

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", errors.New(cfg.InternalError)
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (bool, error) {
	bPubKey, err := ioutil.ReadFile(pubKey)
	if err != nil {
		return false, err
	}

	signingKey, err := jwt.ParseRSAPublicKeyFromPEM(bPubKey)
	if err != nil {
		return false, errors.New(cfg.InternalError)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New(cfg.InvalidToken)
		}
		return signingKey, nil
	})
	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false, errors.New(cfg.InvalidToken)
	}
	if vld, ok := claims["vld"].(bool); !ok || !vld {
		return false, errors.New(cfg.InvalidToken)
	}

	// valid token
	return true, nil
}
