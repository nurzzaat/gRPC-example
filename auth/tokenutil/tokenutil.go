package tokenutil

import (
	"context"
	"fmt"

	"github.com/nurzzaat/gRPC-example/auth/models"

	"errors"
	"strings"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateAccessToken(c context.Context, id uint, secret string, expiry int, redisClient *redis.Client) (accessToken string, err error) {
	//exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := &models.JwtClaims{
		ID: id,
		// RegisteredClaims: jwt.RegisteredClaims{
		// 	ExpiresAt: jwt.NewNumericDate(exp),
		// },
	}
	fmt.Println("claim", claims)
	log.Info("claims:", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	redisClient.Set(c, string(id), t, 0)
	redisToken, err := redisClient.Get(c, string(id)).Result()
	fmt.Println(id, redisToken, err)
	return t, nil
}

func ValidateJWT(c *gin.Context, secret string) error {
	token, err := GetToken(c, secret)
	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func ValidateUserJWT(c *gin.Context, secret string, redisClient *redis.Client) error {
	token, err := GetToken(c, secret)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	userID := uint(claims["id"].(float64))

	var exists bool = false
	redisToken, err := redisClient.Get(c, string(userID)).Result()
	if err != nil {
		return errors.New("invalid token provided")
	}
	if redisToken != "" {
		exists = true
	}

	if ok && token.Valid && exists {
		c.Set("userID", userID)
		return nil
	}
	return errors.New("invalid token provided")
}

func GetToken(c *gin.Context, secret string) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(c)
	if tokenString == "" {
		return nil, errors.New("invalid token provided")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	return token, err
}

func getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
