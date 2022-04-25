package api

import (
	"context"
	"github.com/gin-gonic/gin"
	db "github.com/tornvallalexander/go-backend-template/db/sqlc"
	"github.com/tornvallalexander/go-backend-template/utils"
	"net/http"
	"time"
)

type userResponse struct {
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	CreatedAt         time.Time `json:"created_at"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
}

func newUserResponse(user db.User) *userResponse {
	return &userResponse{
		Username:          user.Username,
		Email:             user.Email,
		CreatedAt:         user.CreatedAt,
		PasswordChangedAt: user.PasswordChangedAt,
	}
}

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=200"`
	Email    string `json:"email" binding:"required,email"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(context.Background(), arg)
	if err != nil {
		status, errRes := checkErr(err)
		ctx.JSON(status, errRes)
		return
	}

	ctx.JSON(http.StatusOK, newUserResponse(user))
}

type getUserRequest struct {
	Username string `uri:"username" binding:"required"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(context.Background(), req.Username)
	if err != nil {
		status, errRes := checkErr(err)
		ctx.JSON(status, errRes)
		return
	}

	ctx.JSON(http.StatusOK, newUserResponse(user))
}

type deleteUserRequest struct {
	Username string `uri:"username" binding:"required"`
}

func (server *Server) deleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.DeleteUser(context.Background(), req.Username)
	if err != nil {
		status, errRes := checkErr(err)
		ctx.JSON(status, errRes)
		return
	}

	ctx.JSON(http.StatusOK, newUserResponse(user))
}
