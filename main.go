package main

import (
	"os"

	"github.com/AkapongAlone/validate-helper/helpers"
	"github.com/AkapongAlone/validate-helper/requests"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	imageFolderPath := os.Getenv("STORAGES_FOLDER_IMAGE_PATH")
	excelFolderPath := os.Getenv("STORAGES_FOLDER_EXCEL_PATH")
	pdfFolderPath := os.Getenv("STORAGES_FOLDER_PDF_PATH")
	htmlFolderPath := os.Getenv("STORAGES_FOLDER_HTML_PATH")
	//allowedOrigins := os.Getenv("ALLOW_ORIGIN")

	r.StaticFS("/storages/excel", gin.Dir(excelFolderPath, true))
	r.StaticFS("/storages/pdf", gin.Dir(pdfFolderPath, true))
	r.StaticFS("/storages/images", gin.Dir(imageFolderPath, true))
	r.StaticFS("/storages/html", gin.Dir(htmlFolderPath, true))
	r.StaticFS("/assets", gin.Dir("./assets", true))
	r.Use(func(c *gin.Context) {

		c.Writer.Header().Set("Content-Type", "application/json")

		// Check if the request's origin is in the list of allowed domains

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	v1 := selectAPIPath(r, os.Getenv("ENV"))
	{
		// Postgres := databases.NewPostgres()

		
		v1.POST("validate", func(c *gin.Context) {
			req := requests.Test{
				ContractNumber: "contract04",
				AcceptList:     "jan",
				TimeNow:        "2022-04-30",
				
			}
			if err := helpers.Validate(req, c); err != nil {
				errResponse := *err
				c.JSON(errResponse.Code, errResponse)
				return
			}

		})
	}

	r.Run(":8080")
}

func selectAPIPath(r *gin.Engine, env string) *gin.RouterGroup {
	if env == "dev" {
		return r.Group("api/v1")
	}

	return r.Group("/v1")
}
