package awsspec

import (
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/hbbio/awless/cloud"
	"github.com/hbbio/awless/logger"
	"github.com/hbbio/awless/template/params"
)

type CreateSubscription struct {
	_        string `action:"create" entity:"subscription" awsAPI:"sns" awsCall:"Subscribe" awsInput:"sns.SubscribeInput" awsOutput:"sns.SubscribeOutput"`
	logger   *logger.Logger
	graph    cloud.GraphAPI
	api      snsiface.SNSAPI
	Topic    *string `awsName:"TopicArn" awsType:"awsstr" templateName:"topic"`
	Endpoint *string `awsName:"Endpoint" awsType:"awsstr" templateName:"endpoint"`
	Protocol *string `awsName:"Protocol" awsType:"awsstr" templateName:"protocol"`
}

func (cmd *CreateSubscription) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("endpoint"), params.Key("protocol"), params.Key("topic")))
}

func (cmd *CreateSubscription) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*sns.SubscribeOutput).SubscriptionArn)
}

type DeleteSubscription struct {
	_      string `action:"delete" entity:"subscription" awsAPI:"sns" awsCall:"Unsubscribe" awsInput:"sns.UnsubscribeInput" awsOutput:"sns.UnsubscribeOutput"`
	logger *logger.Logger
	graph  cloud.GraphAPI
	api    snsiface.SNSAPI
	Id     *string `awsName:"SubscriptionArn" awsType:"awsstr" templateName:"id"`
}

func (cmd *DeleteSubscription) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("id")))
}
