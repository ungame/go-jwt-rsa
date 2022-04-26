package tokens

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"go-jwt-rsa/keypair"
	"log"
	"time"
)

const (
	DefaultExpiration = time.Minute * 10
	DefaultAudience   = "devs"
	DefaultIssuer     = "go-jwt-rsa"
)

func New(id string) (string, error) {

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keypair.GetPrivateKey())
	if err != nil {
		return "", err
	}

	claims := jwt.StandardClaims{
		Audience:  DefaultAudience,
		ExpiresAt: time.Now().Add(DefaultExpiration).Unix(),
		Id:        uuid.New().String(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    DefaultIssuer,
		Subject:   id,
	}

	return newToken(privateKey, claims)
}

func newToken(privateKey *rsa.PrivateKey, claims jwt.StandardClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, &claims)
	return token.SignedString(privateKey)
}

func onParse(token *jwt.Token) (interface{}, error) {

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(keypair.GetPublicKey())
	if err != nil {
		return nil, err
	}

	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("invalid jwt algorithm method: %s", token.Header["alg"])
	}
	return publicKey, nil
}

func Parse(tokenString string) (*jwt.StandardClaims, error) {
	claims := new(jwt.StandardClaims)

	token, err := jwt.ParseWithClaims(tokenString, claims, onParse)
	if err != nil {
		return nil, err
	}

	if !token.Valid || claims == nil {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorMalformed)
	}

	return claims, nil
}

func ToString(token string) string {
	payload, err := Parse(token)

	if err != nil {
		log.Println("error on parse token: ", token)

		return ""
	}

	return fmt.Sprintf(`
----- BEGIN JWT -----

- Token:
	- %s

- Claims:
	- aud: %s
	- exp: %s
	- jti: %s
	- iat: %s
	- iss: %s
	- nbf: %s
	- sub: %s

----- END JWT ------
`,
		token,
		payload.Audience,
		time.Unix(payload.ExpiresAt, 0).String(),
		payload.Id,
		time.Unix(payload.IssuedAt, 0).String(),
		payload.Issuer,
		time.Unix(payload.NotBefore, 0).String(),
		payload.Subject,
	)
}
