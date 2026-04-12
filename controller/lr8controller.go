package controller

import (
	"log/slog"

	"main/logger"

	"github.com/gin-gonic/gin"
)

func HTMLStatic(c *gin.Context) {
	slog.Debug("Отладка", "module", "db")
	slog.Info("Инфо", "port", 8080)
	slog.Warn("Предупреждение", "retry", 3)
	slog.Error("Ошибка", "err", "connection failed")

	logger.GetFileLogger().Debug("message", "123123")

	c.File("../static/html/index.html")
}

func CSSStatic(c *gin.Context) {
	c.File("../static/css/style.css")
}

func JSStatic(c *gin.Context) {
	c.File("../static/js/index.js")
}

func ImageStatic(c *gin.Context) {
	c.File("../static/images/a894b00a1fb5826cbd01aceace20ad06.jpg")
}
