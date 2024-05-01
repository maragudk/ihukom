package html

import (
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents-heroicons/v2/mini"
	"github.com/maragudk/gomponents-heroicons/v2/solid"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"

	"github.com/maragudk/ihukom/model"
)

func HomePage(props PageProps, notes []model.Note) g.Node {
	return page(props,
		Div(ID("notes"),
			//hx.Get("/notes"), hx.Trigger("every 5s"),
			Ol(Class("grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4 text-[8px] sm:text-[10px]"),
				Li(A(hx.Post("/notes"),
					card(
						Div(Class("h-64 cursor-pointer flex justify-center items-center"), solid.Plus(Class("w-12 h-12"))),
					),
				)),
				g.Group(g.Map(notes, func(n model.Note) g.Node {
					return Li(A(Href("/notes?id="+n.ID.String()), hx.Boost("true"),
						card(
							Div(Class("whitespace-pre-wrap font-serif h-64 overflow-hidden"),
								g.Text(n.Content),
							),
						),
					))
				})),
			),
		),
	)
}

func NotePage(props PageProps, n model.Note) g.Node {
	return page(props,
		Div(Class("space-y-4"),
			FormEl(Class("w-full"),
				Input(Type("hidden"), Name("id"), Value(n.ID.String())),
				Textarea(
					ID("autoresize"),
					Name("content"),
					hx.Put("/notes"), hx.Trigger("keyup changed delay:300ms"),
					Placeholder("I think that…"),
					Class("font-serif block w-full bg-white p-4 sm:p-8 shadow rounded-lg border-0 text-gray-900 ring-0 placeholder:text-gray-400 focus:ring-0 sm:text-xl sm:leading-6 overflow-hidden"),
					g.Text(n.Content),
				),
			),

			Div(Class("flex items-center justify-between mt-4 text-orange-600 hover:text-orange-500"),
				A(Href("/"), mini.ArrowLeft(Class("w-5 h-5")), hx.Boost("true")),
				Button(
					mini.Trash(Class("w-5 h-5")),
					hx.Delete("/notes?id="+n.ID.String()), hx.Confirm("Are you sure you want to delete this note?"),
				),
			),
		),
	)
}
