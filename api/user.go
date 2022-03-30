package api

import (
	db "github.com/Adetunjii/simplebank/db/repository"
	"github.com/Adetunjii/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"net/http"
	"time"
)

type createUserRequest struct {
	Username string 	`json:"username" binding:"required,alphanum"`
	Password string 	`json:"password" binding:"required,min=6"`
	FullName string 	`json:"full_name" binding:"required"`
	Email string 		`json:"email" binding:"required,email"`
}

type createUserResponse struct {
	Username string 	`json:"username" binding:"required,alphanum"`
	FullName string 	`json:"full_name" binding:"required"`
	Email string 		`json:"email" binding:"required,email"`
	CreatedAt time.Time 	`json:"created_at"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}


	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserDto{
		Username: req.Username,
		Password: hashedPassword,
		FullName: req.FullName,
		Email: req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {

		if pgxErr, ok := err.(*pgconn.PgError); ok {
			switch pgxErr.Code {
			case pgerrcode.UniqueViolation:
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := createUserResponse{
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}
