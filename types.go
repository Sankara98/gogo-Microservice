package main

import (
	"github.com/cloudnativego/gogo-engine"
)

type matchRepository interface {
	addMatch(match gogo.Match) (err error)
	getMatches() []gogo.Match
	getMatch(id string) (match gogo.Match, err error)
}

type newMatchRequest struct {
	GridSize    int    `json:"gridsize"`
	PlayerWhite string `json:"playerWhite"`
	PlayerBlack string `json:"playerBlack"`
}

func (request newMatchRequest) isValid() (valid bool) {
	valid = true
	if request.GridSize != 19 && request.GridSize != 13 && request.GridSize != 9 {
		valid = false
	}
	if request.PlayerWhite == "" {
		valid = false
	}
	if request.PlayerBlack == "" {
		valid = false
	}
	return valid
}
