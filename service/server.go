package service

import (
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
)

// NewServer configures and returns a Server
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	var matchRespository matchRepository
	mx.HandleFunc("/test", testHandler(formatter)).Methods("Get")
	mx.HandleFunc("/match", createMatchHandler(formatter, matchRespository)).Methods("Post")
}

func testHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK,
			struct{ Test string }{"This is a test"})
	}
}
