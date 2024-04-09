package helper

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var SecretKey = []byte("qwertyuiop")

type Claims struct {
	Email string `json:"username"`
	Role  string `json:"roles"`
	jwt.StandardClaims
}

func JwtTokenStart(c *gin.Context, mail string, role string) {
	_, err := createToken(mail, role)
	if err != nil {
		fmt.Println("")
		fmt.Println("Failed to create Token.......................")
		fmt.Println("")
	}

	// Save the token in a session
	session := sessions.Default(c)
	session.Set(role, mail)
	session.Save()
	check := session.Get(role)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println(check)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}
func createToken(mail string, role string) (string, error) {
	claims := Claims{
		Email: mail,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println(tokenString)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	return tokenString, nil
}

// func ValidateToken(tokenString string, publicKey string) error {
// 	// Parse the token
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Check if the token is signed using the expected algorithm
// 		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
// 			return nil, fmt.Errorf("Err")
// 		}

// 		// Return the public key for verification
// 		return jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
// 	})

// 	// Check if there was an error in parsing the token
// 	if err != nil {
// 		return err
// 	}

// 	// Check if the token is valid
// 	if !token.Valid {
// 		return fmt.Errorf("TokenError")
// 	}

// 	// Check if the token has expired
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
// 		if time.Now().After(expirationTime) {
// 			return fmt.Errorf("TokenExpired")
// 		}
// 	}

// 	// If no errors were found, the token is valid
// 	return nil
// }
