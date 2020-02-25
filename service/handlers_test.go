package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/unrolled/render"
)

const (
	fakeMatchLocationResult = "/matches/5a003b78-409e-4452-b456-a6f0dcee05bd"
)

var (
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
)

func TestCreateMatchRespondsToBadData(t *testing.T) {
	client := &http.Client{}
	var matchRespository matchRepository
	server := httptest.NewServer(http.HandlerFunc(createMatchHandler(formatter, matchRespository)))
	defer server.Close()

	body1 := []byte("this is not valid json")
	body2 := []byte("{\"test\":\"this is valid json, but doesn't conform to server expectations.\"}")

	//Send invalid JSON
	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body1))
	if err != nil {
		t.Errorf("Error in creating POST request for createMatchHandler: %v", err)
	}
	req.Header.Add("Content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in POST to createMatchHandler: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Error("Sending invalid JSON should result in bad request from server")
	}
	req2, err2 := http.NewRequest("POST", server.URL, bytes.NewBuffer(body2))
	if err2 != nil {
		t.Errorf("Error in creating second post request for invalid data on create match: %v", err2)
	}
	req.Header.Add("Content-type", "application/json")
	res2, _ := client.Do(req2)
	defer res2.Body.Close()
	if res2.StatusCode != http.StatusBadRequest {
		t.Error("Sending valid JSON but with incorrect fields should result in a bad request but didn't")
	}
}
func TestCreateMatch(t *testing.T) {
	client := &http.Client{}
	var matchRepository matchRepository
	server := httptest.NewServer(
		http.Handler(createMatchHandler(formatter, matchRepository)))
	defer server.Close()

	body := []byte("{\n \"gridsize\": 19,\n  \"playerWhite\": \"bob\",\n  \"playerBlack\": \"alfred\"\n}")

	req, err := http.NewRequest("POST",
		server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error in creating post request for createMatchHandle: %v", err)
	}

	req.Header.Add("Content", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in POST to createMatchHandler %v", err)
	}

	defer res.Body.Close()
	loc, headerOk := res.Header["Location"]
	if !headerOk {
		t.Errorf("Location header is not set: %v", headerOk)
	} else {
		if !strings.Contains(loc[0], "/matches/") {
			t.Errorf("Location header should contain '/matches/'")
		}
		if len(loc[0]) != len(fakeMatchLocationResult) {
			t.Errorf("Location value does not contain guid of new match")
		}
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected response status 201, received %s", res.Status)
	}
	var matchResponse newMatchResponse
	err = json.Unmarshal(payload, &matchResponse)
	if err != nil {
		t.Errorf("Could not unmarshal payload into newMatchResponse object")
	}

	if matchResponse.ID == "" || !strings.Contains(loc[0], matchResponse.ID) {
		t.Error("matchResponse.Id does not match Location header")
	}
	// After creating a match, match repository should have 1 item in it.
	// matches := matchRepository.getMatches()
	// if err != nil {
	// 	t.Errorf("Unexpected error in getMatches(): %s", err)
	// }
	// if len(matches) != 1 {
	// 	t.Errorf("Expected a match repo of 1 match, got size %d", len(matches))
	// }

	// var match gogo.Match
	// match = matches[0]
	// if match.GridSize != matchResponse.GridSize {
	// 	t.Errorf("Expected repo match and HTTP response gridsize to match. Got %d and %d", match.GridSize, matchResponse.GridSize)
	// }

	// if match.PlayerWhite != "bob" {
	// 	t.Errorf("Repository match, white player should be bob, got %s", match.PlayerWhite)
	// }

	// if matchResponse.PlayerWhite != "bob" {
	// 	t.Errorf("Expected white player to be bob, got %s", matchResponse.PlayerWhite)
	// }

	// if matchResponse.PlayerBlack != "alfred" {
	// 	t.Errorf("Expected black player to be alfred, got %s", matchResponse.PlayerBlack)
	// }
	fmt.Printf("Payload: %s", string(payload))

}
