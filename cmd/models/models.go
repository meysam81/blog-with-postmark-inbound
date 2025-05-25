package models

type InboundEmail struct {
	From        string       `json:"From"`
	FromName    string       `json:"FromName"`
	Subject     string       `json:"Subject"`
	HtmlBody    string       `json:"HtmlBody"`
	TextBody    string       `json:"TextBody"`
	Attachments []Attachment `json:"Attachments"`
}

type Attachment struct {
	Name        string `json:"Name"`
	Content     string `json:"Content"`
	ContentType string `json:"ContentType"`
	ContentID   string `json:"ContentID"`
}

type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	AuthorName  string `json:"author-name"`
	AuthorEmail string `json:"author-email"`
	CreatedAt   string `json:"created-at"`
}

type AuthorizedSender struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created-at"`
}

type EmailInsertDB struct {
	From     string `json:"From"`
	FromName string `json:"FromName"`
	Subject  string `json:"Subject"`
	Content  string `json:"content"`
}
