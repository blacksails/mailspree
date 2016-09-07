package mailspree

// Email represents an email with all the related information.
type Email struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// MailingProvider abstacts different mailing providers.
type MailingProvider interface {
	SendEmail(Email) error
}
