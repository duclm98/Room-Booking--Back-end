package methods

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPassword(hashedPassword string, password string) (isPasswordValid bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}

func CreateToken(userId uint, secretKey string, exprise uint) (string, error) {
  var err error
	token := jwt.New(jwt.SigningMethodHS256)
  claims := token.Claims.(jwt.MapClaims)
  claims["id"] = userId
  claims["exp"] = time.Now().Add(time.Minute * time.Duration(exprise)).Unix()
  t, err := token.SignedString([]byte(secretKey))
  return t, err
}

func ExtractToken(r *http.Request) string {
  bearToken := r.Header.Get("Authorization") //Authorization: bearer token...
  strArr := strings.Split(bearToken, " ")
  if len(strArr) == 2 {
     return strArr[1]
  }
  return ""
}

func VerifyToken(tokenString string, secretKey string) (uint, error) {
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
     //Make sure that the token method conform to "SigningMethodHMAC"
     if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
     }
     return []byte(secretKey), nil
  })
  if err != nil {
    return 0, err
  }
	claims, ok := token.Claims.(jwt.MapClaims)
  if ok && token.Valid {
		idUint64, _ := strconv.ParseUint(fmt.Sprintf("%.f", claims["id"]), 10, 64)
		id := uint(idUint64)
		return id, nil
  }
  return 0, nil
}

func DecodeToken(tokenString string, secretKey string) (uint, error) {
  token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
     //Make sure that the token method conform to "SigningMethodHMAC"
     if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
     }
     return []byte(secretKey), nil
  })
	claims, ok := token.Claims.(jwt.MapClaims)
  if ok && token.Valid {
		idUint64, _ := strconv.ParseUint(fmt.Sprintf("%.f", claims["id"]), 10, 64)
		id := uint(idUint64)
		return id, nil
  }
  return 0, nil
}