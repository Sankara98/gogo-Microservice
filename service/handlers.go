package service

import (
	"encoding/json"
	"github.com/cloudnativego/gogo-engine"
	"github.com/unrolled/render"
	"io/ioutil"
	"net/http"
)

func createMatchHandler(formatter *render.Render, matchRepository matchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		var newMatchRequest newMatchRequest
		err := json.Unmarshal(payload, &newMatchRequest)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest,
				"Failed to parse create match request")
			return
		}
		if !newMatchRequest.isValid() {
			formatter.Text(w, http.StatusBadRequest,
				"Invalid new match request")
			return
		}

		newMatch := gogo.NewMatch(newMatchRequest.GridSize,
			newMatchRequest.PlayerBlack, newMatchRequest.PlayerWhite)
		//matchRepository.addMatch(newMatch)
		w.Header().Add("Location", "/matches/"+newMatch.ID)
		formatter.JSON(w, http.StatusCreated,
			&newMatchResponse{
				ID:          newMatch.ID,
				GridSize:    newMatch.GridSize,
				PlayerWhite: newMatch.PlayerWhite,
				PlayerBlack: newMatch.PlayerBlack})
	}

}
