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
	GetLeague(types.ReqGetDeleteLeague) (types.League, error)
	CreateLeague(types.ReqPostLeague) (string, error)
	UpdateLeague(req types.ReqPutLeague) (types.League, error)
	DeleteLeague(req types.ReqGetDeleteLeague) (string, error)
}

type leagueModel struct {
	db dynamo.DynamoDB
}

func NewLeagueModel(db dynamo.DynamoDB) LeagueModel {
	return &leagueModel{db}
}

func (lm *leagueModel) GetLeague(req types.ReqGetDeleteLeague) (types.League, error) {
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
	item["updated_at"] = now

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

func (lm *leagueModel) UpdateLeague(req types.ReqPutLeague) (types.League, error) {
	res := types.League{}
	svg := lm.db.GetClient()

	// 更新日をjstで作成
	updatedAt, err := attributevalue.Marshal(utils.NowJST())
	if err != nil {
		return res, err
	}

	tableName := os.Getenv("ENV") + "_league"
	updateInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]dynamoTypes.AttributeValue{
			"id": &dynamoTypes.AttributeValueMemberS{Value: req.ID},
		},
		ExpressionAttributeNames: map[string]string{
			"#updated_at": "updated_at",
			"#name":       "name",
			"#manual":     "manual",
		},
		ExpressionAttributeValues: map[string]dynamoTypes.AttributeValue{
			":updated_at": updatedAt,
			":name":       &dynamoTypes.AttributeValueMemberS{Value: req.Name},
			":manual":     &dynamoTypes.AttributeValueMemberS{Value: *req.Manual},
		},
		UpdateExpression: aws.String("SET #updated_at = :updated_at,#name = :name,#manual = :manual"),
		ReturnValues:     dynamoTypes.ReturnValueAllNew,
	}

	result, err := svg.UpdateItem(context.TODO(), updateInput)
	if err != nil {
		return res, err
	}

	err = attributevalue.UnmarshalMap(result.Attributes, &res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (lm *leagueModel) DeleteLeague(req types.ReqGetDeleteLeague) (string, error) {
	svg := lm.db.GetClient()

	tableName := os.Getenv("ENV") + "_league"
	deleteInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]dynamoTypes.AttributeValue{
			"id": &dynamoTypes.AttributeValueMemberS{Value: req.ID},
		},
	}

	_, err := svg.DeleteItem(context.TODO(), deleteInput)
	if err != nil {
		return "", err
	}

	return req.ID, err
}
