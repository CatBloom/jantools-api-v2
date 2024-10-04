package types

import "time"

type ReqGetGameList struct {
	LeagueID string `query:"leagueID" validate:"required"`
}

type ReqGetDeleteGame struct {
	ID       string `query:"id" validate:"required"`
	LeagueID string `query:"leagueID" validate:"required"`
}

type ReqPostGame struct {
	LeagueID string   `json:"leagueID" validate:"required" dynamodbav:"league_id"`
	Results  []Result `json:"results" validate:"dive,required" dynamodbav:"results" `
}

type ReqPutGame struct {
	ID       string   `json:"id" validate:"required" dynamodbav:"id"`
	LeagueID string   `json:"leagueID" validate:"required" dynamodbav:"league_id"`
	Results  []Result `json:"results" validate:"dive,required" dynamodbav:"results" `
}

type Game struct {
	ID        string    `json:"id" dynamodbav:"id"`
	LeagueID  string    `json:"leagueID" dynamodbav:"league_id"`
	CreatedAt time.Time `json:"createdAt" dynamodbav:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" dynamodbav:"updated_at"`
	Results   []Result  `json:"results"  dynamodbav:"results" `
}

type Result struct {
	Rank      int    `json:"rank" validate:"required" dynamodbav:"rank"`
	Name      string `json:"name" validate:"required" dynamodbav:"name"`
	Point     int    `json:"point" validate:"required" dynamodbav:"point"`
	CalcPoint int    `json:"calcPoint" validate:"required" dynamodbav:"calcPoint"`
}
