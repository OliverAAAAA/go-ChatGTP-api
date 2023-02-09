package handlers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-chatgpt-api/config"
	"go-chatgpt-api/openai"
	services "go-chatgpt-api/service"
	"net/http"
)

var askLogService services.AskLogService

func Ask(c *gin.Context) {
	secret := c.Query("secret")

	if secret == "" || secret != *config.GetRequestSecret() {
		c.JSON(500, gin.H{
			"error": "秘钥错误，拒绝访问",
			"code":  "500",
		})
		return
	}
	content := c.Query("content")
	result, err := requestOpenAIChat(c, content)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg":  "OpenAI官方系统繁忙，请稍后再试",
			"code": "500",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": result,
		"msg":  "success",
		"code": "200",
	})
}

func requestOpenAIChat(c *gin.Context, content string) (string, error) {
	result, err := openai.Ask(content, c.ClientIP())
	if err != nil {
		return "", err
	}
	return result, nil
}

func requestOpenAICreateImg(c *gin.Context, content string) (string, error) {
	result, err := openai.GenerateImg(content, c.ClientIP())
	if err != nil {
		return "", err
	}
	return result, nil
}

func AskSearch(c *gin.Context) {
	content := c.PostForm("content")
	var result string
	if content != "" {
		var err error
		result, err = requestOpenAIChat(c, content)
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusOK, "search.html", gin.H{
				"data":    "OpenAI官方系统繁忙，请稍后再试",
				"content": content,
			})
			return
		}
	}
	c.HTML(http.StatusOK, "search.html", gin.H{
		"data":    result,
		"content": content,
	})
}

func CreateImg(c *gin.Context) {
	content := c.PostForm("createImgMsg")
	var url string
	if content != "" {
		var err error
		url, err = requestOpenAICreateImg(c, content)
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusOK, "search.html", gin.H{
				"data":         "OpenAI官方系统繁忙，请稍后再试",
				"createImgMsg": content,
			})
			return
		}
	}
	c.HTML(http.StatusOK, "search.html", gin.H{
		"url":          url,
		"createImgMsg": content,
	})
}

func AskStream(c *gin.Context) {
	content := c.Query("content")
	openai.AskStream(content, c.ClientIP())
	//if err != nil {
	//	log.Println(err)
	//	c.JSON(500, gin.H{
	//		"error": "系统繁忙",
	//	})
	//	return
	//}
	c.JSON(200, gin.H{
		"data": 1,
	})
}

func GenerateImg(c *gin.Context) {
	content := c.Query("content")
	url, err := requestOpenAICreateImg(c, content)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": "OpenAI官方系统繁忙，请稍后再试",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": url,
	})
}
