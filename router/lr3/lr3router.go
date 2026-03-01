package lr3

import (
	"main/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouterForLR3(router *gin.RouterGroup) {
	router.GET("/html", controller.HtmlController)

	router.GET("/text", controller.TextController)

	router.GET("/json", controller.JsonController)

	router.GET("/xml", controller.XmlController)

	router.GET("/csv", controller.CsvController)

	router.GET("/binary", controller.BinaryController)

	router.GET("/image", controller.ImageController)

	router.GET("/pdf", controller.PdfController)

	router.GET("/redirect301", controller.Redirect301Controller)

	router.GET("/redirect302", controller.Redirect302Controller)
}
