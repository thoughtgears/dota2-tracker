package dota

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ProMatch struct {
	MatchId       int64  `json:"match_id"`
	Duration      int    `json:"duration"`
	StartTime     int    `json:"start_time"`
	RadiantTeamId int    `json:"radiant_team_id"`
	RadiantName   string `json:"radiant_name"`
	DireTeamId    int    `json:"dire_team_id"`
	DireName      string `json:"dire_name"`
	Leagueid      int    `json:"leagueid"`
	LeagueName    string `json:"league_name"`
	SeriesId      int    `json:"series_id"`
	SeriesType    int    `json:"series_type"`
	RadiantScore  int    `json:"radiant_score"`
	DireScore     int    `json:"dire_score"`
	RadiantWin    bool   `json:"radiant_win"`
	Radiant       bool   `json:"radiant"`
}

func (c *Client) GetProMatches() ([]ProMatch, error) {
	var matches []ProMatch
	req, _ := http.NewRequest("GET", c.baseUrl+"/proMatches", nil)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get pro matches: %w", err)
	}

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get pro matches - %v: %s", string(body), resp.Status)
	}

	if err := json.Unmarshal(body, &matches); err != nil {
		return nil, fmt.Errorf("failed to unmarshal pro matches: %w", err)
	}

	return matches, nil
}
