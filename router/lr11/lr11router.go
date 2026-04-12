package lr11

import (
	"main/controller"
	"main/database/repos"

	"github.com/gin-gonic/gin"
)

func RegisterRouterForLR11(router *gin.RouterGroup, ur *repos.UserRepo, br *repos.BalanceRepo, tr *repos.TransactionRepo) {
	userController := controller.CreateUserController(ur, tr)
	balanceController := controller.CreateBalanceController(br, tr)

	router.GET("/user/:id", userController.GetUser)
	router.POST("/user", userController.CreateUser)
	router.PATCH("/user/:id", userController.UpdateUser)
	router.DELETE("/user/:id", userController.DeleteUser)

	router.POST("/create/:userId", balanceController.CreateBalance)
	router.POST("/add", balanceController.AddBalance)
	router.POST("/subtract", balanceController.SubtractBalance)
}
