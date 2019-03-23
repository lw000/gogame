package ggmail

import (
	"errors"
	"gopkg.in/gomail.v2"
)

func SendMail(cfg *MailConfig, subject string, body string) error {
	if cfg == nil {
		return errors.New("MailConfig is nil")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "tuyue"+"<"+cfg.From+">")
	//收件人
	m.SetHeader("To", cfg.To...)
	//抄送
	m.SetHeader("Cc", cfg.From)
	////暗送
	//m.SetHeader("BCc",  cfg.From)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	//m.Attach("这是附件")
	d := gomail.NewDialer(cfg.Host, int(cfg.Port), cfg.From, cfg.Pass)
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}
