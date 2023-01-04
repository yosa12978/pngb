package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yosa12978/pngb/internal/pkg/dtos"
	"github.com/yosa12978/pngb/internal/pkg/helpers"
	"github.com/yosa12978/pngb/internal/pkg/repositories"
	"github.com/yosa12978/pngb/internal/pkg/services"
)

func NewHandler() http.Handler {
	postService := services.NewPostService(repositories.NewPostMongoRepo())
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		posts := postService.Find()
		helpers.RenderTemplate(w, "./templates/index.html", posts)
	}).Methods("GET")

	router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		post, err := postService.FindByID(vars["id"])
		if err != nil {
			http.Error(w, "not found", 404)
			return
		}
		helpers.RenderTemplate(w, "./templates/detail.html", post)
	}).Methods("GET")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		dto := dtos.PostCreateDTO{
			Text: r.FormValue("Text"),
			Imgs: []string{r.FormValue("Img1")},
		}
		post, err := dto.Map()
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if err = postService.Create(post); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		http.Redirect(w, r, "/", 301)
	}).Methods("POST")

	return router
}
