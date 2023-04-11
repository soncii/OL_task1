package transport

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (h *Handler) jwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing token")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return h.JWTSecret, nil
		})
		fmt.Println(err)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		email := claims["sub"].(string)

		c.Set("email", email)

		return next(c)
	}
}

func (h *Handler) GenerateJWT(userEmail string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userEmail,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(h.JWTSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
