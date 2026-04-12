package controller

import (
	"strconv"

	"main/database/repos"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UR *repos.UserRepo
	TR *repos.TransactionRepo
}

func CreateUserController(ur *repos.UserRepo, tr *repos.TransactionRepo) *UserController {
	return &UserController{UR: ur, TR: tr}
}

func (ub *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	logs, _ := ub.TR.GetLogsByUserID(id)
	ctx.JSON(200, gin.H{
		"result": gin.H{
			"userData": ub.UR.GetUser(id),
			"userLogs": logs,
		},
	})
}

func (ub *UserController) CreateUser(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"result": ub.UR.CreateUser(user.Login),
	})
}

func (ub *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"result": ub.UR.DeleteUser(id, "login"),
	})
}

func (ub *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var user User
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"result": ub.UR.UpdateUser(id, user.Login),
	})
}

// ---------------------
//
//

type BalanceController struct {
	BR *repos.BalanceRepo
	TR *repos.TransactionRepo
}

func CreateBalanceController(repo *repos.BalanceRepo, tr *repos.TransactionRepo) *BalanceController {
	return &BalanceController{BR: repo, TR: tr}
}

func (c *BalanceController) CreateBalance(ctx *gin.Context) {
	userID, err := strconv.ParseInt(ctx.Param("userId"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid user id"})
		return
	}

	balance, err := c.BR.CreateBalance(userID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, balance)
}

func (c *BalanceController) AddBalance(ctx *gin.Context) {
	var req struct {
		UserID int64 `json:"user_id"`
		Amount int64 `json:"amount"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	balance, err := c.BR.AddBalance(req.UserID, req.Amount)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.TR.CreateLog(balance.ID, req.Amount)

	ctx.JSON(200, balance)
}

func (c *BalanceController) SubtractBalance(ctx *gin.Context) {
	var req struct {
		UserID int64 `json:"user_id"`
		Amount int64 `json:"amount"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	balance, err := c.BR.SubtractBalance(req.UserID, req.Amount)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.TR.CreateLog(balance.ID, req.Amount)

	ctx.JSON(200, balance)
}
