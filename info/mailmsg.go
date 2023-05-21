package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"golang.org/x/exp/slog"
	"net/smtp"
)

type mailMsg struct {
	Level string
	Msg   string
}

func sendMsgMail(c *gin.Context) {
	m := mailMsg{}
	if err := c.BindJSON(&m); err != nil {
		slog.Error("解析json失败", "msg", err)
		Error(c, "解析消息失败")
		return
	}
	switch m.Level {
	case "Info":
		if err := msgMail(m.Msg); err != nil {
			Error(c, "发送失败")
			return
		}
	case "Error":
		if err := msgMail("不要灰心勇于尝试，你定会成功！\n" + m.Msg); err != nil {
			Error(c, "发送失败")
			return
		}
	default:
		Error(c, "未知等级，不发送")
		return
	}
	OK(c)
}
func msgMail(msg string) error {
	e := &email.Email{
		From:    "2025907338@qq.com",
		To:      []string{"2025907338@qq.com"},
		Subject: "进程通知邮件",
	}
	e.Text = []byte(msg)
	if err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "2025907338@qq.com", "**", "smtp.qq.com")); err != nil {
		slog.Error("发送邮件失败", "msg", err)
		return err
	}
	slog.Info("发送邮件成功!")
	return nil
}
