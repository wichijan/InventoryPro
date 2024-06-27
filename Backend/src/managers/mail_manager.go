package managers

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
)

type MailMgr interface {
	SendWelcomeMail(to string, username string) *models.INVError
	// SendOrderConfirmationMail(to string, order models.GetOrderDTO) *models.KTSError
	SendLinkForNewPasswordMail(to string, userId *string) *models.INVError
}

type MailManager struct {
	MailgunInstance utils.MailgunInterface
}

const emailSender = "Cinemika Team <team@cinemika.tech>"

func (mm *MailManager) SendWelcomeMail(to string, username string) *models.INVError {
	subject := "Welcome to Cinemika!"

	body, err := utils.PrepareWelcomeMailBody(username)
	if err != nil {
		return inv_errors.INV_UPSTREAM_ERROR
	}

	return mm.sendMail(to, emailSender, subject, body)
}

/*
func (mm *MailManager) SendOrderConfirmationMail(to string, order models.GetOrderDTO) *models.KTSError {
	subject := "Order confirmation"

	body, err := utils.PrepareOrderConfirmationBody(order)
	if err != nil {
		return inv_errors.KTS_UPSTREAM_ERROR
	}

	return mm.sendMail(to, emailSender, subject, body)
}
*/

func (mm *MailManager) sendMail(to string, from, subject, body string) *models.INVError {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	message := mm.MailgunInstance.NewMessage(emailSender, subject, "", to)
	message.AddHeader("Content-Type", "text/html")
	message.SetHtml(body)

	_, _, err := mm.MailgunInstance.Send(ctx, message)
	if err != nil {
		log.Println("Error in MailManager.SendMail().MailgunInstance.Send(): ", err.Error())
		return inv_errors.INV_UPSTREAM_ERROR
	}

	return nil
}

func InitializeMailgunClient() *mailgun.MailgunImpl {
	ApiKey := os.Getenv("MAILGUN_API_KEY")
	Domain := os.Getenv("MAILGUN_DOMAIN")

	log.Println("Initializing Mailgun client...")
	log.Println("Domain: ", Domain)

	mg := mailgun.NewMailgun(Domain, ApiKey)
	mg.SetAPIBase("https://api.mailgun.net/v3")

	return mg
}

func (mm *MailManager) SendLinkForNewPasswordMail(to string, userId *string) *models.INVError {
	subject := "Änder dein Passwort"

	body, err := utils.PrepareResetPasswordBody(userId)
	if err != nil {
		return inv_errors.INV_UPSTREAM_ERROR
	}

	return mm.sendMail(to, emailSender, subject, body)
}
