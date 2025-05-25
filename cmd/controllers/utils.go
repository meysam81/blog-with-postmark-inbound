package routes

import (
	"strings"

	. "github.com/meysam81/tarzan/cmd/models"
)

func maskEmail(p *Post) {
	at := strings.Index(p.AuthorEmail, "@")
	if at <= 1 {
		p.AuthorEmail = "***"
		return
	}
	name := p.AuthorEmail[:at]
	domain := p.AuthorEmail[at:]
	if len(name) <= 2 {
		p.AuthorEmail = string(name[0]) + "*" + domain
		return
	}
	p.AuthorEmail = string(name[0]) + strings.Repeat("*", len(name)-2) + string(name[len(name)-1]) + domain
}
