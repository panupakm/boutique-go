package biz

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/panupakm/boutique-go/api/email"
	"github.com/panupakm/boutique-go/app/email/internal/conf"
)

type mailInfo struct {
	subject  string
	template *template.Template
}

type mailInfoCenter struct {
	orderConfirmation mailInfo
}

type EmailUseCase struct {
	log            *log.Helper
	mailInfoCenter mailInfoCenter
	from           string
	addr           string
	auth           smtp.Auth
	mailCenter     mailInfoCenter
}

func createMapOfTemplates(dir string) (map[string]*template.Template, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	entriesMap := make(map[string]*template.Template, len(entries))
	for _, entry := range entries {
		entriesMap[entry.Name()] = template.Must(template.ParseFiles(dir + "/" + entry.Name()))
	}

	return entriesMap, nil
}

func NewEmailUseCase(conf *conf.Data, logger log.Logger) *EmailUseCase {
	server := conf.GetMailServer()
	from := server.GetFrom()
	addr := fmt.Sprintf("%s:%d", server.GetHost(), server.GetPort())

	templates, err := createMapOfTemplates(server.GetTemplateDir())
	if err != nil {
		panic(err)
	}

	confirmation := mailInfo{
		subject:  conf.GetOrderConfirmMail().GetSubject(),
		template: templates[conf.GetOrderConfirmMail().GetTemplateName()],
	}

	return &EmailUseCase{
		from: from,
		addr: addr,
		auth: smtp.PlainAuth("", addr, server.GetPassword(), addr),
		mailCenter: mailInfoCenter{
			orderConfirmation: confirmation,
		},
	}
}

func buildEmailContent(tmpl *template.Template, subject string, to string, from string, data interface{}) string {
	var bodyBuf bytes.Buffer
	tmpl.Execute(&bodyBuf, data)
	return fmt.Sprintf("To: %s\n", to) +
		fmt.Sprintf("Subject: %s\n", subject) +
		fmt.Sprint("MIME-version: 1.0;\n") +
		fmt.Sprint("Content-type: text/html; charset=\"UTF-8\";\n\n") +
		bodyBuf.String()
}

func (ec *EmailUseCase) SendOrderConfirmation(ctx context.Context, to string, data interface{}) error {
	mailInfo := &ec.mailCenter.orderConfirmation
	msg := buildEmailContent(mailInfo.template, mailInfo.subject, to, ec.from, data)
	err := smtp.SendMail(ec.addr, nil, ec.from, []string{to}, []byte(msg))
	if err != nil {
		ec.log.Error(err)
		return email.ErrorConnectionLost("EmailService: %s", err.Error())
	}
	return nil
}
