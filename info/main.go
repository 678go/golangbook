package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/info", sendMsgMail)
	r.Run(":60123")
}
