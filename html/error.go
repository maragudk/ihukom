package html

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func ErrorPage() g.Node {
	return page(PageProps{Title: "Something went wrong", Description: "Oh no! 😵"},
		prose(
			H1(g.Text("Something went wrong")),
			P(g.Text("Oh no! 😵")),
			P(A(Href("/"), g.Text("Back to front."))),
		),
	)
}
