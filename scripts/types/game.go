package types

import "time"

type ReqGame struct {
	ID       string `query:"id" validate:"required"`
	LeagueID string `query:"league_id" validate:"required"`
}

type Game struct {
	ID        string    `json:"id" dynamodbav:"id"`
	League_ID string    `json:"leagueID" dynamodbav:"league_id"`
	CreatedAt time.Time `json:"createdAt" dynamodbav:"created_at"`
	Results   []Result  `json:"results" dynamodbav:"results"`
}

type Result struct {
	Rank  int    `json:"rank" dynamodbav:"rank"`
	Name  string `json:"name" dynamodbav:"name"`
	Point int    `json:"point" dynamodbav:"point"`
}
