package response

import "github.com/gin-gonic/gin"

func OK(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"status": "success",
	})
}

func SendMessage(c *gin.Context, message string) {
	c.IndentedJSON(200, gin.H{
		"status":  "success",
		"message": message,
	})
}

func SendData(c *gin.Context, data interface{}) {
	c.IndentedJSON(200, gin.H{
		"status": "success",
		"data":   data,
	})
}
