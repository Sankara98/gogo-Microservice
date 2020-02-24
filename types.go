package main

import (
	"github.com/cloudnativego/gogo-engine"
)

type matchRepository interface {
	addMatch(match gogo.Match) (err error)
	getMatches() []gogo.Match
	getMatch(id string) (match gogo.Match, err error)
}
