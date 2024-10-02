package types

import "time"

type ReqGameList struct {
	LeagueID string `query:"leagueID" validate:"required"`
}

type ReqGame struct {
	ID       string `query:"id" validate:"required"`
	LeagueID string `query:"leagueID" validate:"required"`
}

type Game struct {
	ID        string    `json:"id" dynamodbav:"id"`
	League_ID string    `json:"leagueID"  validate:"required" dynamodbav:"league_id"`
	CreatedAt time.Time `json:"createdAt" dynamodbav:"created_at"`
	Results   []Result  `json:"results" validate:"dive,required" dynamodbav:"results" `
}

type Result struct {
	Rank      int    `json:"rank" validate:"required" dynamodbav:"rank"`
	Name      string `json:"name" validate:"required" dynamodbav:"name"`
	Point     int    `json:"point" validate:"required" dynamodbav:"point"`
	CalcPoint int    `json:"calcPoint" validate:"required" dynamodbav:"calcPoint"`
}
