package user

import (
	"encoding/json"
	"errors"
	. "github.com/VigneshSK17/cubimer-api/src/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
	"net/http"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type JWTClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "request does not contain access token", http.StatusUnauthorized)
			return
		}

		err := validateToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GenerateToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request CreateUserParams

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := Instance.Queries.GetUserByUsername(ctx, request.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = checkPassword(user, request.Password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tokenString, err := generateJWT(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{"token": tokenString})
}

func generateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &JWTClaim{
		username,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func validateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return errors.New("token expired")
	}

	return nil
}
