package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	"github.com/matcornic/hermes/v2"
)

type MailgunInterface interface {
	Send(ctx context.Context, message *mailgun.Message) (mes string, id string, err error)
	NewMessage(from, subject, text string, to ...string) *mailgun.Message
}

var h = hermes.Hermes{
	Product: hermes.Product{
		Name:        "InventoryPro",
		Link:        URL,
		TroubleText: "Wenn der {ACTION}-Button nicht funktioniert, kopieren Sie einfach die URL unten und fügen Sie sie in Ihren Webbrowser ein.",
		Copyright:   "Copyright © 2024 InventoryPro-WWI22SEB",
	},
	Theme: new(hermes.Default),
}

func PrepareWelcomeMailBody(username string) (string, error) {
	hermesMail := hermes.Email{
		Body: hermes.Body{
			Name: username,
			Intros: []string{
				fmt.Sprintf("Willkommen bei InventoryPro, %v! Wir freuen uns dich an Board zu haben .", username),
			},
			Outros: []string{
				"Wenn du Hilfe brauchst, oder Fragen hast, antworte einfach auf diese E-Mail, wir helfen dir gerne weiter.",
			},
		},
	}

	body, err := h.GenerateHTML(hermesMail)

	if err != nil {
		return "", err
	}

	return body, nil
}

func prettierTime(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

func prettyPrice(price int32) string {
	x := float64(price)
	x = x / 100
	return fmt.Sprintf("€%.2f", x)
}

/*
func generateDataFromOrder(order models.GetOrderDTO) [][]hermes.Entry {
	var data [][]hermes.Entry
	for _, ticket := range order.Tickets {
		data = append(data, []hermes.Entry{
			{Key: "Price Category", Value: ticket.PriceCategory.CategoryName},
			{Key: "Seat Category", Value: ticket.SeatCategory.CategoryName},
			{Key: "Row", Value: fmt.Sprint(ticket.Seat.RowNr)},
			{Key: "Column", Value: fmt.Sprint(ticket.Seat.ColumnNr)},
		})
	}
	return data
}
*/
/*
func PrepareOrderConfirmationBody(order models.GetOrderDTO) (string, error) {

	hermesMail := hermes.Email{
		Body: hermes.Body{
			Intros: []string{
				"Your order has been processed successfully.",
			},
			Dictionary: []hermes.Entry{
				{Key: "Event Title", Value: order.Event.Title},
				{Key: "Cinema Hall", Value: order.CinemaHall.Name},
				{Key: "Date and Time", Value: prettierTime(order.Event.Start)},
				{Key: "Total Price", Value: prettyPrice(order.Order.Totalprice)},
				{Key: "Theatre Name", Value: order.Theatre.Name},
			},
			Table: hermes.Table{

				Data: generateDataFromOrder(order),

				Columns: hermes.Columns{
					CustomWidth: map[string]string{
						"Price Category": "35%",
						"Seat Category":  "35%",
						"Row":            "15%",
						"Column":         "15%",
					},
				},
			},
			Actions: []hermes.Action{
				{
					Instructions: "You can check your order and more in your overview:",
					Button: hermes.Button{
						Text:  "Go to Overview",
						Link:  "http://" + URL + "/overview",
						Color: "#334155",
					},
				},
			},
		},
	}

	body, err := h.GenerateHTML(hermesMail)

	if err != nil {
		return "", err
	}

	return body, nil
}
*/

func PrepareResetPasswordBody(userId *string) (string, error) {
	hermesMail := hermes.Email{
		Body: hermes.Body{
			Intros: []string{
				"Du hast eine Anfrage zur Passwortänderung gestellt.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Klick den Button, um dein Passwort zu ändern:",
					Button: hermes.Button{
						Text:  "Änder dein Passwort",
						Link:  "http://" + URL + "/reset-password?userId=" + *userId,
						Color: "#334155",
					},
				},
			},
			Outros: []string{
				"Wenn du keine Passwortänderung angefordert haben, ignoriere diese E-Mail bitte.",
			},
		},
	}

	body, err := h.GenerateHTML(hermesMail)

	if err != nil {
		return "", err
	}

	return body, nil
}

func PrepareInformAdminsRegistBody(username string) (string, error) {
	hermesMail := hermes.Email{
		Body: hermes.Body{
			Intros: []string{
				"Benutzer " + username + " hat sich registriert.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Klick den Button, um den Benutzer anzunehmen oder abzulehnen:",
					Button: hermes.Button{
						Text:  "anzunehmen oder abzulehnen",
						Link:  "http://" + URL + "/admin/users",
						Color: "#334155",
					},
				},
			},
		},
	}

	body, err := h.GenerateHTML(hermesMail)

	if err != nil {
		return "", err
	}

	return body, nil
}



func PrepareInformUserItemAcceptBody() (string, error) {
	hermesMail := hermes.Email{
		Body: hermes.Body{
			Intros: []string{
				"Item-Transfer wurde angenommen",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Kontrolliere den Transfer in deinem Dashboard:",
					Button: hermes.Button{
						Text:  "Check Transfer",
						Link:  "http://" + URL + "/dashboard",
						Color: "#334155",
					},
				},
			},
		},
	}

	body, err := h.GenerateHTML(hermesMail)

	if err != nil {
		return "", err
	}

	return body, nil
}


func PrepareInformUserItemRequestBody() (string, error) {
	hermesMail := hermes.Email{
		Body: hermes.Body{
			Intros: []string{
				"Anfrage auf Item-Transfer.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Jemand möchte dir sein Item transferieren:",
					Button: hermes.Button{
						Text:  "Anzunehmen oder abzulehnen",
						Link:  "http://" + URL + "/dashboard",
						Color: "#334155",
					},
				},
			},
		},
	}

	body, err := h.GenerateHTML(hermesMail)

	if err != nil {
		return "", err
	}

	return body, nil
}