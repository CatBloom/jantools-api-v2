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

type GameModel interface {
	GetGameList(req types.ReqGetGameList) ([]types.Game, error)
	GetGame(types.ReqGetDeleteGame) (types.Game, error)
	CreateGame(types.ReqPostGame) (string, error)
	UpdateGame(req types.ReqPutGame) (types.Game, error)
	DeleteGame(req types.ReqGetDeleteGame) (string, error)
}

type gameModel struct {
	db dynamo.DynamoDB
}

func NewGameModel(db dynamo.DynamoDB) GameModel {
	return &gameModel{db}
}

func (gm *gameModel) GetGameList(req types.ReqGetGameList) ([]types.Game, error) {
	res := []types.Game{}
	svg := gm.db.GetClient()

	tableName := os.Getenv("ENV") + "_game"
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("league_id = :leagueID"),
		ExpressionAttributeValues: map[string]dynamoTypes.AttributeValue{
			":leagueID": &dynamoTypes.AttributeValueMemberS{Value: req.LeagueID},
		},
		ConsistentRead: aws.Bool(true),
	}

	result, err := svg.Query(context.TODO(), queryInput)
	if err != nil {
		return res, err
	}

	if result.Items == nil {
		return res, errors.New("dynamodb item not found")
	}

	err = attributevalue.UnmarshalListOfMaps(result.Items, &res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (gm *gameModel) GetGame(req types.ReqGetDeleteGame) (types.Game, error) {
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

func (gm *gameModel) CreateGame(req types.ReqPostGame) (string, error) {
	svg := gm.db.GetClient()

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

func (gm *gameModel) UpdateGame(req types.ReqPutGame) (types.Game, error) {
	res := types.Game{}
	svg := gm.db.GetClient()

	// 更新日をjstで作成
	updatedAt, err := attributevalue.Marshal(utils.NowJST())
	if err != nil {
		return res, err
	}

	resultList, err := attributevalue.MarshalList(req.Results)
	if err != nil {
		return res, err
	}

	tableName := os.Getenv("ENV") + "_game"
	updateInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]dynamoTypes.AttributeValue{
			"id":        &dynamoTypes.AttributeValueMemberS{Value: req.ID},
			"league_id": &dynamoTypes.AttributeValueMemberS{Value: req.LeagueID},
		},
		ExpressionAttributeNames: map[string]string{
			"#results":    "results",
			"#updated_at": "updated_at",
		},
		ExpressionAttributeValues: map[string]dynamoTypes.AttributeValue{
			":results":    &dynamoTypes.AttributeValueMemberL{Value: resultList},
			":updated_at": updatedAt,
		},
		UpdateExpression: aws.String("SET #results = :results,#updated_at = :updated_at"),
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

func (gm *gameModel) DeleteGame(req types.ReqGetDeleteGame) (string, error) {
	svg := gm.db.GetClient()

	tableName := os.Getenv("ENV") + "_game"
	deleteInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]dynamoTypes.AttributeValue{
			"id":        &dynamoTypes.AttributeValueMemberS{Value: req.ID},
			"league_id": &dynamoTypes.AttributeValueMemberS{Value: req.LeagueID},
		},
	}

	_, err := svg.DeleteItem(context.TODO(), deleteInput)
	if err != nil {
		return "", err
	}

	return req.ID, err
}
