package api

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	db "github.com/JairoRiver/registro_gastos/tree/main/Backend/internal/db/sqlc"
	"github.com/JairoRiver/registro_gastos/tree/main/Backend/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (server *Server) getUserFromToken(ctx *gin.Context) (db.GetUserByUsernameRow, error) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	user, err := server.store.GetUserByUsername(context.Background(), authPayload.Username)
	if err != nil {

		return user, err
	}
	return user, nil
}

// createBook handler
type createEntryRequest struct {
	TypeID        string  `json:"type_id" binding:"required,uuid"`
	Name          string  `json:"name" binding:"required"`
	UseDay        int64   `json:"use_day" binding:"required min=0"`
	Amount        float64 `json:"amount" binding:"required"`
	Cost          float64 `json:"cost" binding:"required"`
	CostIndicator string  `json:"cost_indicator" binding:"omitempty"`
	Place         string  `json:"place" binding:"required"`
}

func (server *Server) createEntry(ctx *gin.Context) {
	user, err := server.getUserFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req createEntryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	typeID, err := uuid.Parse(req.TypeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateEntryParams{
		UserID: user.ID,
		TypeID: typeID,
		Name:   req.Name,
		UseDay: time.Unix(req.UseDay, 0).Local(),
		Amount: req.Amount,
		Cost:   req.Cost,
		Place:  req.Place,
	}

	//Validate is cost indicator is valid
	if len(req.CostIndicator) > 0 {
		arg.CostIndicator = sql.NullString{String: req.CostIndicator, Valid: true}
	}

	entry, err := server.store.CreateEntry(ctx, arg)
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
	ctx.JSON(http.StatusOK, entry)
}

// getEntry handler
type getEntryRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) getEntry(ctx *gin.Context) {
	user, err := server.getUserFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req getEntryRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	entryID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	entry, err := server.store.GetEntry(ctx, entryID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if entry.UserID != user.ID {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entry)
}

// listEntry by User handler
func (server *Server) listEntries(ctx *gin.Context) {
	var req getEntryRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, err := server.getUserFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if userID != user.ID {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	entries, err := server.store.ListEntryByUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entries)
}

// updateEntry handler
type updateEntryRequestData struct {
	TypeID        string          `json:"type_id" binding:"omitempty,uuid"`
	Name          string          `json:"name" binding:"omitempty"`
	UseDay        int64           `json:"use_day" binding:"omitempty"`
	Amount        sql.NullFloat64 `json:"amount" binding:"omitempty"`
	Cost          sql.NullFloat64 `json:"cost" binding:"omitempty"`
	CostIndicator sql.NullString  `json:"cost_indicator" binding:"omitempty"`
	Place         string          `json:"place" binding:"omitempty"`
}
type updateEntryRequestID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) updateEntry(ctx *gin.Context) {
	var entryIDReq updateEntryRequestID
	if err := ctx.ShouldBindUri(&entryIDReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	entryID, err := uuid.Parse(entryIDReq.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	oldEntry, err := server.store.GetEntry(ctx, entryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, err := server.getUserFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if oldEntry.UserID != user.ID {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	var reqData updateEntryRequestData
	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateEntryParams{
		ID: entryID,
	}

	//Validate is the typeID is valid
	if len(reqData.TypeID) > 0 {
		typeID, err := uuid.Parse(reqData.TypeID)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}

		type_model, err := server.store.GetType(ctx, typeID)

		if err != nil {
			if err == sql.ErrNoRows {
				err = errors.New("entry need a valid type")
				ctx.JSON(http.StatusBadRequest, errorResponse(err))
				return
			}

			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		arg.TypeID = uuid.NullUUID{UUID: type_model.ID, Valid: true}
	}

	//Validate is the name is valid
	if len(reqData.Name) > 0 {
		arg.Name = sql.NullString{String: reqData.Name, Valid: true}
	}

	//Validate is the use_day is valid
	if reqData.UseDay > 0 {
		arg.UseDay = sql.NullTime{Time: time.Unix(reqData.UseDay, 0).Local(), Valid: true}
	}

	//Validate is the amount is valid
	if reqData.Amount.Valid {
		arg.Amount = reqData.Amount
	}

	//Validate is the cost is valid
	if reqData.Cost.Valid {
		arg.Cost = reqData.Cost
	}

	//Validate is the Cost Indicator is valid
	if reqData.CostIndicator.Valid {
		arg.CostIndicator = reqData.CostIndicator
	}

	//Validate is the place is valid
	if len(reqData.Place) > 0 {
		arg.Place = sql.NullString{String: reqData.Place, Valid: true}
	}

	entry, err := server.store.UpdateEntry(ctx, arg)

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

	ctx.JSON(http.StatusOK, entry)
}

// delete Entry handler
type deleteEntryRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) deleteEntry(ctx *gin.Context) {
	var req deleteEntryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	entryID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	entry, err := server.store.GetEntry(ctx, entryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, err := server.getUserFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if entry.UserID != user.ID {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	err = server.store.DeleteEntry(ctx, entryID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entryID)
}
