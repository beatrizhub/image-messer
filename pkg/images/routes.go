package images

import (
	"github.com/go-chi/chi"
)

func (s *Server) Routes() chi.Router {

	r := chi.NewRouter()

	r.Get("/", s.showUploadForm)
	r.Post("/upload", s.uploadImage)
	r.Get("/download/{filename}", s.downloadImage)

	return r

}