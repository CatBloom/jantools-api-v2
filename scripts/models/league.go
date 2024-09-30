package models

import (
	"context"
	"errors"
	"main/dynamo"
	customTypes "main/types"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type LeagueModel interface {
	GetLeagueByID(string) (customTypes.League, error)
	CreateLeague(customTypes.League) (string, error)
}

type leagueModel struct {
	db dynamo.DynamoDB
}

func NewLeagueModel(db dynamo.DynamoDB) LeagueModel {
	return &leagueModel{db}
}

func (lm *leagueModel) GetLeagueByID(id string) (customTypes.League, error) {
	res := customTypes.League{}
	svg := lm.db.GetClient()

	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String("league"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
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

func (lm *leagueModel) CreateLeague(req customTypes.League) (string, error) {
	svg := lm.db.GetClient()

	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	// uuidからハイフンを除去
	id := strings.ReplaceAll(uuid.String(), "-", "")
	req.ID = id

	// 作成日をjstで作成
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	req.CreatedAt = time.Now().In(jst)

	item, err := attributevalue.MarshalMap(req)
	if err != nil {
		return "", err
	}

	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String("league"),
		Item:      item,
	}

	_, err = svg.PutItem(context.TODO(), putItemInput)
	if err != nil {
		return "", err
	}

	return id, err
}
