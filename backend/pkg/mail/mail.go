package mail

type MailResponse struct {
	Summary string
	Mails   []Mail
}
type Mail struct {
	From    string
	Date    string
	Subject string
	Body    string
	//Images  []Image
	//Links   []Link
	Images  []string
	Links   []string
	Summary string
}

type Image struct {
	Source      string
	Description string
}

type Link struct {
	HRef        string
	Description string
}
