package types

import "time"

type ReqGetGameList struct {
	LeagueID string `query:"leagueID" validate:"required"`
}

type ReqGetGame struct {
	ID       string `query:"id" validate:"required"`
	LeagueID string `query:"leagueID" validate:"required"`
}

type ReqPostGame struct {
	LeagueID string    `json:"leagueID" dynamodbav:"league_id"`
	GameDate time.Time `json:"gameDate" validate:"required" dynamodbav:"game_date"`
	Results  []Result  `json:"results" validate:"dive,required" dynamodbav:"results" `
}

type ReqPutGame struct {
	ID       string    `json:"id" validate:"required" dynamodbav:"id"`
	LeagueID string    `json:"leagueID" dynamodbav:"league_id"`
	GameDate time.Time `json:"gameDate" validate:"required" dynamodbav:"game_date"`
	Results  []Result  `json:"results" validate:"dive,required" dynamodbav:"results" `
}

type ReqDeleteGame struct {
	ID       string `query:"id" validate:"required"`
	LeagueID string `query:"leagueID"`
}

type Game struct {
	ID        string    `json:"id" dynamodbav:"id"`
	LeagueID  string    `json:"leagueID" dynamodbav:"league_id"`
	CreatedAt time.Time `json:"createdAt" dynamodbav:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" dynamodbav:"updated_at"`
	GameDate  time.Time `json:"gameDate" dynamodbav:"game_date"`
	Results   []Result  `json:"results" dynamodbav:"results" `
}

type Result struct {
	Rank      int      `json:"rank" validate:"required" dynamodbav:"rank"`
	Name      string   `json:"name" validate:"required" dynamodbav:"name"`
	Point     *int     `json:"point" validate:"required" dynamodbav:"point"`
	CalcPoint *float64 `json:"calcPoint" validate:"required" dynamodbav:"calcPoint"`
}
