package awsspec

import (
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/hbbio/awless/cloud"
	"github.com/hbbio/awless/logger"
	"github.com/hbbio/awless/template/params"
)

type CreateTopic struct {
	_      string `action:"create" entity:"topic" awsAPI:"sns" awsCall:"CreateTopic" awsInput:"sns.CreateTopicInput" awsOutput:"sns.CreateTopicOutput"`
	logger *logger.Logger
	graph  cloud.GraphAPI
	api    snsiface.SNSAPI
	Name   *string `awsName:"Name" awsType:"awsstr" templateName:"name"`
}

func (cmd *CreateTopic) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("name")))
}

func (cmd *CreateTopic) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*sns.CreateTopicOutput).TopicArn)
}

type DeleteTopic struct {
	_      string `action:"delete" entity:"topic" awsAPI:"sns" awsCall:"DeleteTopic" awsInput:"sns.DeleteTopicInput" awsOutput:"sns.DeleteTopicOutput"`
	logger *logger.Logger
	graph  cloud.GraphAPI
	api    snsiface.SNSAPI
	Id     *string `awsName:"TopicArn" awsType:"awsstr" templateName:"id"`
}

func (cmd *DeleteTopic) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("id")))
}
