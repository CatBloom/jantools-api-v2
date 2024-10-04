package models

import (
	"context"
	"errors"
	"main/dynamo"
	"main/types"
	"main/utils"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamoTypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type LeagueModel interface {
	GetLeague(types.ReqGetLeague) (types.League, error)
	CreateLeague(types.ReqPostLeague) (string, error)
}

type leagueModel struct {
	db dynamo.DynamoDB
}

func NewLeagueModel(db dynamo.DynamoDB) LeagueModel {
	return &leagueModel{db}
}

func (lm *leagueModel) GetLeague(req types.ReqGetLeague) (types.League, error) {
	res := types.League{}
	svg := lm.db.GetClient()

	tableName := os.Getenv("ENV") + "_league"
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]dynamoTypes.AttributeValue{
			"id": &dynamoTypes.AttributeValueMemberS{Value: req.ID},
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

func (lm *leagueModel) CreateLeague(req types.ReqPostLeague) (string, error) {
	svg := lm.db.GetClient()

	item, err := attributevalue.MarshalMap(req)
	if err != nil {
		return "", err
	}

	id, err := utils.GenerateUUIDWithoutHyphens()
	if err != nil {
		return "", err
	}
	// // 作成日、更新日をjstで作成
	now, err := attributevalue.Marshal(utils.NowJST())
	if err != nil {
		return "", err
	}
	item["id"] = &dynamoTypes.AttributeValueMemberS{Value: id}
	item["created_at"] = now

	tableName := os.Getenv("ENV") + "_league"
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
