package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int
	Msg  string
}

func OK(c *gin.Context) {
	c.JSON(http.StatusOK, &Response{
		Code: 200,
		Msg:  "success",
	})
}

func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, &Response{
		Code: -1,
		Msg:  msg,
	})
}
