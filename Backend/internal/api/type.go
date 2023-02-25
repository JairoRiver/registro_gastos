package api

import (
	"database/sql"
	"net/http"

	db "github.com/JairoRiver/registro_gastos/tree/main/Backend/internal/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// createType handler
type createTypeRequest struct {
	Name string `json:"name" binding:"required,alphanum"`
}

func (server *Server) createType(ctx *gin.Context) {
	var req createTypeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	type_model, err := server.store.CreateType(ctx, req.Name)
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

	ctx.JSON(http.StatusOK, type_model)
}

// getType handler
type getTypeRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) getType(ctx *gin.Context) {
	var req getTypeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	typeID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	type_model, err := server.store.GetType(ctx, typeID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, type_model)
}

// listTypes handler
func (server *Server) listTypes(ctx *gin.Context) {

	types_model, err := server.store.ListTypes(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, types_model)
}

// updateType handler
type updateTypeRequestData struct {
	Name string `json:"name" binding:"omitempty,alphanum"`
}
type updateTypeRequestID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) updateType(ctx *gin.Context) {
	var typeIDReq updateTypeRequestID
	if err := ctx.ShouldBindUri(&typeIDReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	typeID, err := uuid.Parse(typeIDReq.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	var reqData updateTypeRequestData
	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateTypeParams{
		ID: typeID,
	}

	//Validate is the name is valid
	if len(reqData.Name) > 0 {
		arg.Name = sql.NullString{String: reqData.Name, Valid: true}
	}

	type_model, err := server.store.UpdateType(ctx, arg)

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

	ctx.JSON(http.StatusOK, type_model)

}

// deleteType handler
type deleteTypeRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) deleteType(ctx *gin.Context) {
	var req deleteTypeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	typeID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteType(ctx, typeID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, typeID)
}
