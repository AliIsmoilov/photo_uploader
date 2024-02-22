package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// 	AKIARH6EDLOEQPAQ244U
	//	ZxzrALIv1VFee6W2R2pBHCVzehM2/GmzVxzdRF71

	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.POST("/", func(c *gin.Context) {

		// Get the file
		file, err := c.FormFile("image")
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "Failed to upload image",
			})
			return
		}

		err = c.SaveUploadedFile(file, "assets/uploads"+file.Filename)
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "Failed to upload image",
			})
			return
		}

		// Render the page
		c.HTML(http.StatusOK, "index.html", gin.H{
			"image": "assets/uploads" + file.Filename,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
