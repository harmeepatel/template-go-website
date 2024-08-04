package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"makeyourcareer.net/web/pages"
)

func Index() http.Handler {
	return templ.Handler(pages.Index("LINK1"))
}

func Page2() http.Handler {
	return templ.Handler(pages.ExtraPage("LINK2", "LINK2"))
}

func Page3() http.Handler {
	return templ.Handler(pages.ExtraPage("LINK3", "LINK3"))
}
