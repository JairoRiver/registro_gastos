package api

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/JairoRiver/registro_gastos/tree/main/Backend/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// adminUserMiddleware creates a gin middleware for validate is the user is valid
type getID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

var UserErr = errors.New("user unauthorized")

func userMiddleware(server *Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

		user, err := server.store.GetUserByUsername(context.Background(), authPayload.Username)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.AbortWithStatusJSON(http.StatusNotFound, errorResponse(err))
				return
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		var req getID
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		path := strings.Split(ctx.FullPath(), "/")
		endpointType := path[2]

		switch endpointType {
		case "user":
			userID, err := uuid.Parse(req.ID)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

			if userID == user.ID {
				ctx.Next()
				return
			}

		case "entry", "entries", "group", "groups":
			entryID, err := uuid.Parse(req.ID)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

			entry, err := server.store.GetEntry(context.Background(), entryID)
			if err != nil {
				if err == sql.ErrNoRows {
					ctx.AbortWithStatusJSON(http.StatusNotFound, errorResponse(err))
					return
				}

				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

			if entry.UserID == user.ID {
				ctx.Next()
				return
			}

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(UserErr))
		}
	}
}
