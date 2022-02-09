package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		text := c.PostForm("text")
		lang := c.DefaultPostForm("lang", "ZH")

		client := resty.New()

		resp, err := client.R().
			SetHeader("Content-Type", "application/x-www-form-urlencoded").
			SetFormData(map[string]string{
				"auth_key":    "key",
				"text":        text,
				"target_lang": "ZH",
			}).
			Post("https://api.deepl.com/v2/translate")

		if err != nil {
			c.String(http.StatusOK, fmt.Sprintln(resp))
		} else {
			c.String(http.StatusOK, fmt.Sprintln(resp))
		}

	})

	err := router.Run(":8080")

	if err != nil {
		fmt.Println("err:", err)
	}
}
