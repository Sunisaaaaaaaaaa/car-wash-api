package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
		if header := ctx.Request.Header.Get("Authorization"); header != "" {
			tokenString := strings.Replace(header, "Bearer ", "", 1)
			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
				}
				return hmacSampleSecret, nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx.Set("userId", claims["userId"])
				ctx.Set("role", claims["role"])
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"errorCode":   http.StatusInternalServerError,
					"errorDetail": err.Error(),
				})
				return
			}
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"errorCode":   http.StatusInternalServerError,
				"errorDetail": "token not found",
			})
			return
		}
		ctx.Next()
	}
}

func Authorization(validRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Request.Header.Get("Authorization")) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"errorCode":   http.StatusUnauthorized,
				"errorDetail": "Unauthorized access",
			})
			return
		}

		if roleVal, ok := ctx.Get("role"); !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"errorCode":   http.StatusUnauthorized,
				"errorDetail": "role is not set",
			})
			return
		} else {
			var roles []int
			if v, ok := roleVal.(float64); ok {
				roles = []int{int(v)}
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"errorCode":   http.StatusInternalServerError,
					"errorDetail": "role invalid",
				})
				return
			}

			valid := make(map[int]int)
			for _, val := range roles {
				valid[val] = 0
			}

			// at least one role is in validRoles
			for _, valStr := range validRoles {
				val, err := strconv.Atoi(valStr)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				if _, ok := valid[val]; ok {
					ctx.Next()
					return
				}
			}
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"errorCode":   http.StatusUnauthorized,
			"errorDetail": "You are not authorized to access this path",
		})

	}
}

func LogoutHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if header := ctx.Request.Header.Get("Authorization"); header != "" {
			ctx.Request.Header.Del("Authorization")
			ctx.JSON(http.StatusOK, gin.H{
				"message": "logout successful",
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"errorCode":   http.StatusInternalServerError,
				"errorDetail": "token not found",
			})
		}
	}
}
