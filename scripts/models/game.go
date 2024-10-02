package models

import (
	"context"
	"errors"
	"main/dynamo"
	"main/types"
	"main/utils"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamoTypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type GameModel interface {
	GetGame(types.ReqGame) (types.Game, error)
	CreateGame(types.Game) (string, error)
}

type gameModel struct {
	db dynamo.DynamoDB
}

func NewGameModel(db dynamo.DynamoDB) GameModel {
	return &gameModel{db}
}

func (gm *gameModel) GetGame(req types.ReqGame) (types.Game, error) {
	res := types.Game{}
	svg := gm.db.GetClient()

	tableName := os.Getenv("ENV") + "_game"
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]dynamoTypes.AttributeValue{
			"id":        &dynamoTypes.AttributeValueMemberS{Value: req.ID},
			"league_id": &dynamoTypes.AttributeValueMemberS{Value: req.LeagueID},
		},
		ConsistentRead: aws.Bool(true),
	}

	result, err := svg.GetItem(context.TODO(), getItemInput)
	if err != nil {
		return res, err
	}

	if result.Item == nil {
		return res, errors.New("dynamodb item not found")
	}

	err = attributevalue.UnmarshalMap(result.Item, &res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (gm *gameModel) CreateGame(req types.Game) (string, error) {
	svg := gm.db.GetClient()

	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	// uuidからハイフンを除去
	id := strings.ReplaceAll(uuid.String(), "-", "")
	req.ID = id

	// 作成日をjstで作成
	req.CreatedAt = utils.NowJST()

	item, err := attributevalue.MarshalMap(req)
	if err != nil {
		return "", err
	}

	tableName := os.Getenv("ENV") + "_game"
	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	}

	_, err = svg.PutItem(context.TODO(), putItemInput)
	if err != nil {
		return "", err
	}

	return id, err
}
