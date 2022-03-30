package api

import (
	"database/sql"
	"fmt"
	db "github.com/Adetunjii/simplebank/db/repository"
	"github.com/Adetunjii/simplebank/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type transferDto struct {
	SourceAccountID int64 	`json:"source_account_id" binding:"required,min=1"`
	DestinationAccountID int64 	`json:"destination_account_id" binding:"required,min=1"`
	Amount int64 				`json:"amount" binding:"required,gt=0"`
	Currency string				`json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferDto
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validAccount(ctx, req.SourceAccountID, req.Currency) {
		return
	}

	if !server.validAccount(ctx, req.DestinationAccountID, req.Currency) {
		return
	}

	arg := db.TransferTxnParams{
		SourceAccountID:      req.SourceAccountID,
		DestinationAccountID: req.DestinationAccountID,
		Amount:               req.Amount,
		Currency:             req.Currency,
		Reference:            util.RandomString(10),
	}

	account, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string)bool {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account %d::  %s vs %s", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}


	return true
 }