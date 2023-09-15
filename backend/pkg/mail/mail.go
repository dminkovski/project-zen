package mail

type MailResponse struct {
	Summary string
	Mails   []*Mail
}
type Mail struct {
	From    string
	Date    string
	Subject string
	Body    string
	//Images  []Image
	//Links   []Link
	Images  []string `json:"Images,omitempty"`
	Links   []string `json:"Links,omitempty"`
	Summary string   `json:"Summary,omitempty"`
}

type Image struct {
	Source      string
	Description string
}

type Link struct {
	HRef        string
	Description string
}
