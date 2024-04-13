package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/config"
	"github.com/isaqueveras/powersso/domain/authentication"
	"github.com/isaqueveras/powersso/tokens"
)

// Session models the session data
type Session struct {
	SessionID uuid.UUID
	UserID    uuid.UUID
	UserLevel string
	FirstName string
}

// GetSession gets the session data from the context
func GetSession(ctx *gin.Context) *Session {
	session, ok := ctx.Get("SESSION")
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return nil
	}

	value := session.(jwt.MapClaims)
	return &Session{
		SessionID: uuid.MustParse(value["SessionID"].(string)),
		UserID:    uuid.MustParse(value["UserID"].(string)),
		UserLevel: value["UserLevel"].(string),
		FirstName: value["FirstName"].(string),
	}
}

// Auth is a middleware to check if the user is authorized to access the resource
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		if token = ctx.GetHeader("Authorization"); token == "" || len(token) < 30 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims := tokens.ParseJWT(token[7:], config.Get().GetSecrets()); claims != nil {
			ctx.Set("UID", claims["UserID"])
			ctx.Set("SESSION", claims)
			return
		}

		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}

// OnlyAdmin check if the user is an administrator
func OnlyAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if GetSession(ctx).UserLevel != string(authentication.AdminLevel) {
			session := GetSession(ctx)
			log.Printf("WARNING: user (%v - %v) tried to access user tried to access route for administrators only", session.UserID, session.FirstName)
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
		ctx.Next()
	}
}

// Yourself validates if the logged in user is the same as the request
func Yourself() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := GetSession(ctx)
		userIn := ctx.Param("user_id")

		if session.UserID.String() != userIn {
			log.Printf("WARNING: user (%v - %v) tried to access information for user (%v)", session.UserID, session.FirstName, userIn)
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}
