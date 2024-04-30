package html

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func HomePage(props PageProps) g.Node {
	return page(props,
		H1(Class("text-2xl font-bold text-gray-900 sm:text-4xl"), g.Text(`ihukom`)),
	)
}
