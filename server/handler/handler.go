package handler

import (
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		err := c.SaveUploadedFile(file, "images/"+file.Filename)

	}
}
