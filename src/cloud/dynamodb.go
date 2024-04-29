/*
 * @Author: charley zhu
 * @Date: 2023-10-12 12:15:42
 * @LastEditTime: 2023-10-12 14:30:32
 * @LastEditors: charley zhu
 * @Description:
 */
package cloud

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"go.uber.org/zap"

	"global-backend/src/logger"
)

var sess = new(session.Session)
var dynamoDbClient = new(dynamodb.DynamoDB)

func Scan(tableName string) ([]interface{}, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	ret, err := dynamoDbClient.Scan(input)
	if err != nil {
		logger.Logger.Error("scan table failed", zap.Error(err))
		return nil, err
	}
	if ret.Items == nil {
		logger.Logger.Error("scan table failed, no item")
		return nil, err
	}
	items := make([]interface{}, 0)
	err = dynamodbattribute.UnmarshalListOfMaps(ret.Items, &items)
	if err != nil {
		logger.Logger.Error("scan table failed", zap.Error(err))
		return nil, err
	}
	return items, nil
}

func Get(tableName, key string) (interface{}, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(key),
			},
		},
	}
	ret, err := dynamoDbClient.GetItem(input)
	if err != nil {
		logger.Logger.Error("get item failed", zap.Error(err))
		return nil, err
	}
	if ret.Item == nil {
		logger.Logger.Error("get item failed, no item")
		return nil, err
	}
	item := make(map[string]interface{})
	err = dynamodbattribute.UnmarshalMap(ret.Item, &item)
	if err != nil {
		logger.Logger.Error("get item failed", zap.Error(err))
		return nil, err
	}
	return item, nil
}

func init() {
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           "env_stg",
	}))

	dynamoDbClient = dynamodb.New(sess)
}
