package types

import (
	"time"
)

type League struct {
	ID        string    `json:"id" dynamodbav:"id"`
	CreatedAt time.Time `json:"created_at" dynamodbav:"created_at"`
	Name      string    `json:"name" validate:"required" dynamodbav:"name"`
	Manual    string    `json:"manual" dynamodbav:"manual"`
	StartAt   string    `json:"startAt" dynamodbav:"start_at"`
	FinishAt  string    `json:"finishAt" dynamodbav:"finish_at"`
	Rule      Rule      `json:"rule" dynamodbav:"rule"`
	Players   []Player  `json:"players" dynamodbav:"players"`
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

type Player struct {
	ID   int    `json:"id" validate:"required" dynamodbav:"id"`
	Name string `json:"name" validate:"required" dynamodbav:"name"`
}
