package utils

import (
	"github.com/L4B0MB4/Musicfriends/pkg/models"
	"github.com/L4B0MB4/Musicfriends/pkg/server/interfaces"
	"github.com/gin-gonic/gin"
)

func GetUserContextFromCtx(ctx *gin.Context, sessionStore interfaces.SessionStore) *models.UserContext {

	userKey := ctx.GetString("session")
	userContext, _ := sessionStore.GetSession(userKey)
	return &userContext
}
