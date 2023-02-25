package api

import (
	"database/sql"
	"net/http"

	db "github.com/JairoRiver/registro_gastos/tree/main/Backend/internal/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// createGroup handler
type createGroupRequest struct {
	Name   string `json:"name" binding:"required,alphanum"`
	UserID string `json:"user_id" binding:"required,uuid"`
}

func (server *Server) createGroup(ctx *gin.Context) {
	var req createGroupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateGroupParams{
		UserID: userID,
		Name:   req.Name,
	}

	group, err := server.store.CreateGroup(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, group)
}

// getGroup handler
type getGroupRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) getGroup(ctx *gin.Context) {
	var req getGroupRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	groupID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	group, err := server.store.GetGroup(ctx, groupID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, group)
}

// listgroups handler
type listGroupsRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) listGroups(ctx *gin.Context) {
	var req listGroupsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	groups, err := server.store.ListGroupsByUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, groups)
}

// updateGroup handler
type updateGroupRequestData struct {
	Name string `json:"username" binding:"omitempty,alphanum"`
}
type updateGroupRequestID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) updateGroup(ctx *gin.Context) {
	var groupIDReq updateGroupRequestID
	if err := ctx.ShouldBindUri(&groupIDReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	groupID, err := uuid.Parse(groupIDReq.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	var reqData updateGroupRequestData
	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateGroupsParams{
		ID: groupID,
	}

	//Validate is the name is valid
	if len(reqData.Name) > 0 {
		arg.Name = sql.NullString{String: reqData.Name, Valid: true}
	}

	group, err := server.store.UpdateGroups(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, group)

}

// deleteGroup handler
type deleteGroupRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) deleteGroup(ctx *gin.Context) {
	var req deleteGroupRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	groupID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteGroup(ctx, groupID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, groupID)
}
