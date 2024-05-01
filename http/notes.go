package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	g "github.com/maragudk/gomponents"
	hxhttp "github.com/maragudk/gomponents-htmx/http"
	ghttp "github.com/maragudk/gomponents/http"
	"github.com/maragudk/httph"

	"github.com/maragudk/ihukom/html"
	"github.com/maragudk/ihukom/model"
)

type NoteRequest struct {
	ID      model.ID
	Content string
}

type noteGetter interface {
	GetNotes(ctx context.Context) ([]model.Note, error)
}

func Home(mux chi.Router, db noteGetter) {
	mux.Get("/", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
		notes, err := db.GetNotes(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return html.ErrorPage(), err
		}

		if hxhttp.IsRequest(r.Header) && !hxhttp.IsBoosted(r.Header) {
			return html.NotesPartial(notes), nil
		}

		return html.HomePage(html.PageProps{Title: "Home", Description: ""}, notes), nil
	}))
}

type noteGetSaver interface {
	GetNote(ctx context.Context, id model.ID) (model.Note, error)
	CreateNote(ctx context.Context) (model.Note, error)
	SaveNote(ctx context.Context, n model.Note) error
	DeleteNote(ctx context.Context, id model.ID) error
}

func Notes(mux chi.Router, db noteGetSaver) {
	mux.Get("/notes", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return nil, nil
		}

		n, err := db.GetNote(r.Context(), model.ID(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return html.ErrorPage(), err
		}

		return html.NotePage(html.PageProps{Title: ""}, n), nil
	}))

	mux.Post("/notes", func(w http.ResponseWriter, r *http.Request) {
		n, err := db.CreateNote(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		hxhttp.SetRedirect(w.Header(), "/notes?id="+n.ID.String())
	})

	mux.Put("/notes", httph.FormHandler(func(w http.ResponseWriter, r *http.Request, req NoteRequest) {
		n := model.Note{ID: req.ID, Content: req.Content}
		if err := db.SaveNote(r.Context(), n); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))

	mux.Delete("/notes", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if err := db.DeleteNote(r.Context(), model.ID(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		hxhttp.SetRedirect(w.Header(), "/")
	})
}
