package utils

import (
	"context"
	"fmt"

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
			Name:      username,
			Signature: "Mit freundlichen Grüßen",
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

func PrepareResetPasswordBody(userId *string) (string, error) {
	hermesMail := hermes.Email{
		Body: hermes.Body{
			Signature: "Mit freundlichen Grüßen",
			Intros: []string{
				"Du hast eine Anfrage zur Passwortänderung gestellt.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Klick den Button, um dein Passwort zu ändern:",
					Button: hermes.Button{
						Text:  "Änder dein Passwort",
						Link:  "http://" + URL + "/auth/forgot-password/" + *userId,
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
			Signature: "Mit freundlichen Grüßen",
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
			Signature: "Mit freundlichen Grüßen",
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
			Signature: "Mit freundlichen Grüßen",
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

func PrepareRegistrationCodeBody(code string) (string, error) {
	hermesMail := hermes.Email{
		Body: hermes.Body{
			Signature: "Mit freundlichen Grüßen",
			Intros: []string{
				"Anmeldung-Abschließen.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Um die Registrierung abzuschließen, nutze diesen Einmal-Code um ein neues Password festzulegen:",
					InviteCode:   code,
					Button: hermes.Button{
						Text:  "Code verwenden",
						Link:  "http://" + URL + "/auth/code/code?=" + code,
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
