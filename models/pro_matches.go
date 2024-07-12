package models

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
