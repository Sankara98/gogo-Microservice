package service

import (
	// "fmt"
	"net/http"

	// "github.com/cloudfoundry-community/go-cfenv"
	// cftools "github.com/cloudnativego/cf-tools"
	// "github.com/cloudnativego/cfmgo"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

// func initRepository() (repo matchRepository) {
// 	appEnv, _ := cfenv.Current()
// 	dbServiceURI, err := cftools.GetVCAPServiceProperty(dbServiceName, "url", appEnv)
// 	if err != nil || dbServiceURI == "" {
// 		if err != nil {
// 			fmt.Printf("\nError retrieving database configuration: %v\n", err)
// 		}
// 		fmt.Println("MongoDB was not detected; configuring inMemoryRepository...")
// 		repo = newInMemoryRepository()
// 		return
// 	}
// 	matchCollection := cfmgo.Connect(cfmgo.NewCollectionDialer, dbServiceURI,
// 		MatchesCollectionName)
// 	fmt.Printf("Connecting to MongoDB service: %s...\n", dbServiceName)
// 	repo = newMongoMatchRepository(matchCollection)
// 	return

// }

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
