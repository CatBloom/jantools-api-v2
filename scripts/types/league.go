package types

import (
	"time"
)

type ReqGetLeague struct {
	ID string `query:"id" validate:"required"`
}

type ReqPostLeague struct {
	Name     string `json:"name" validate:"required" dynamodbav:"name"`
	Manual   string `json:"manual" dynamodbav:"manual"`
	Password string `json:"password" dynamodbav:"password"`
	Rule     Rule   `json:"rule" dynamodbav:"rule"`
}

type ReqPutLeague struct {
	ID     string  `json:"id" dynamodbav:"id" `
	Name   string  `json:"name" validate:"required" dynamodbav:"name"`
	Manual *string `json:"manual" validate:"required" dynamodbav:"manual"`
}

type ReqDeleteLeague struct {
	ID string `query:"id"`
}

type League struct {
	ID        string    `json:"id" dynamodbav:"id"`
	CreatedAt time.Time `json:"createdAt" dynamodbav:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" dynamodbav:"updated_at"`
	Name      string    `json:"name" dynamodbav:"name"`
	Manual    string    `json:"manual" dynamodbav:"manual"`
	Password  string    `json:"password" dynamodbav:"password"`
	Rule      Rule      `json:"rule" dynamodbav:"rule"`
}

type Rule struct {
	PlayerCount int    `json:"playerCount" validate:"required" dynamodbav:"player_count"`
	GameType    string `json:"gameType" validate:"required" dynamodbav:"game_type"`
	Tanyao      *bool  `json:"tanyao" dynamodbav:"tanyao"`
	Back        *bool  `json:"back" validate:"required" dynamodbav:"back"`
	Dora        *int   `json:"dora" validate:"required" dynamodbav:"dora"`
	StartPoint  *int   `json:"startPoint" validate:"required" dynamodbav:"start_point"`
	ReturnPoint *int   `json:"returnPoint" validate:"required" dynamodbav:"return_point"`
	Uma         []int  `json:"uma" validate:"required" dynamodbav:"uma"`
}
