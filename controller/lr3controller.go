package controller

import (
	"encoding/xml"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Person struct {
	XMLName   xml.Name `xml:"person" json:"-"`
	ID        int      `json:"id" xml:"id"`
	FirstName string   `json:"first_name" xml:"firstName"`
	LastName  string   `json:"last_name" xml:"lastName"`
	Age       int      `json:"age" xml:"age"`
	Email     string   `json:"email" xml:"email"`
}

type Company struct {
	XMLName     xml.Name `xml:"company"`
	Name        string   `xml:"name"`
	FoundedYear int      `xml:"foundedYear"`
	Employees   int      `xml:"employees"`
	Address     string   `xml:"address"`
}

func HtmlController(c *gin.Context) {
	htmlContent := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>HTML ответ</title>
	</head>
	<body>
		<h1>HTML ответ от Gin сервера</h1>
	</body>
	</html>`

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
}

func TextController(c *gin.Context) {
	text := "Простой текстовый ответ от TextController\n"
	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(text))
}

func JsonController(c *gin.Context) {
	people := []Person{
		{ID: 1, FirstName: "Иван", LastName: "Петров", Age: 30, Email: "ivan@example.com"},
		{ID: 2, FirstName: "Мария", LastName: "Иванова", Age: 25, Email: "maria@example.com"},
		{ID: 3, FirstName: "Петр", LastName: "Сидоров", Age: 35, Email: "petr@example.com"},
	}

	c.JSON(http.StatusOK, people)
}

func XmlController(c *gin.Context) {
	company := Company{
		Name:        "Example Corp",
		FoundedYear: 2020,
		Employees:   150,
		Address:     "ул. Примерная, 123, Москва",
	}

	c.XML(http.StatusOK, company)
}

func CsvController(c *gin.Context) {
	bytes, err := os.ReadFile("./sources/people-100.csv")
	if err != nil {
		panic(err.Error())
	}
	c.Data(http.StatusOK, "text/csv; charset=utf-8", bytes)
}

func BinaryController(c *gin.Context) {
	bytes, err := os.ReadFile("./sources/a894b00a1fb5826cbd01aceace20ad06.jpg")
	if err != nil {
		panic(err.Error())
	}
	c.Data(http.StatusOK, "application/octet-stream", bytes)
}

func ImageController(c *gin.Context) {
	bytes, err := os.ReadFile("./sources/a894b00a1fb5826cbd01aceace20ad06.jpg")
	if err != nil {
		panic(err.Error())
	}
	c.Data(http.StatusOK, "image/png", bytes)
}

func PdfController(c *gin.Context) {
	bytes, err := os.ReadFile("./sources/ЛР2_Журбей_А.М._241-334.pdf")
	if err != nil {
		panic(err.Error())
	}

	c.Data(http.StatusOK, "application/pdf", bytes)
}

func Redirect301Controller(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/lr3/html")
}

func Redirect302Controller(c *gin.Context) {
	c.Redirect(http.StatusFound, "/lr3/text")
}
