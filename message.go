package mailspree

import "bytes"

// Message represents an email with all the related information.
type Message struct {
	From    Email   `json:"from"`
	To      []Email `json:"to"`
	Subject string  `json:"subject" valid:"required"`
	Body    string  `json:"body" valid:"required"`
}

// Email is simply a real name together with the address
type Email struct {
	Name    string `json:"name"`
	Address string `json:"address" valid:"email,required"`
}

// String returns a string representation of an Email
func (e Email) String() string {
	b := bytes.NewBufferString(e.Name)
	if e.Name != "" {
		b.WriteString(" <")
		b.WriteString(e.Address)
		b.WriteString(">")
	} else {
		b.WriteString(e.Address)
	}
	return b.String()
}
