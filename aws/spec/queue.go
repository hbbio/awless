package awsspec

import (
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/hbbio/awless/cloud"
	"github.com/hbbio/awless/logger"
	"github.com/hbbio/awless/template/params"
)

type CreateQueue struct {
	_                 string `action:"create" entity:"queue" awsAPI:"sqs" awsCall:"CreateQueue" awsInput:"sqs.CreateQueueInput" awsOutput:"sqs.CreateQueueOutput"`
	logger            *logger.Logger
	graph             cloud.GraphAPI
	api               sqsiface.SQSAPI
	Name              *string `awsName:"QueueName" awsType:"awsstr" templateName:"name"`
	Delay             *string `awsName:"Attributes[DelaySeconds]" awsType:"awsstringpointermap" templateName:"delay"`
	MaxMsgSize        *string `awsName:"Attributes[MaximumMessageSize]" awsType:"awsstringpointermap" templateName:"max-msg-size"`
	RetentionPeriod   *string `awsName:"Attributes[MessageRetentionPeriod]" awsType:"awsstringpointermap" templateName:"retention-period"`
	Policy            *string `awsName:"Attributes[Policy]" awsType:"awsstringpointermap" templateName:"policy"`
	MsgWait           *string `awsName:"Attributes[ReceiveMessageWaitTimeSeconds]" awsType:"awsstringpointermap" templateName:"msg-wait"`
	RedrivePolicy     *string `awsName:"Attributes[RedrivePolicy]" awsType:"awsstringpointermap" templateName:"redrive-policy"`
	VisibilityTimeout *string `awsName:"Attributes[VisibilityTimeout]" awsType:"awsstringpointermap" templateName:"visibility-timeout"`
}

func (cmd *CreateQueue) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("name"),
		params.Opt("delay", "max-msg-size", "msg-wait", "policy", "redrive-policy", "retention-period", "visibility-timeout"),
	))
}

func (cmd *CreateQueue) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*sqs.CreateQueueOutput).QueueUrl)
}

type DeleteQueue struct {
	_      string `action:"delete" entity:"queue" awsAPI:"sqs" awsCall:"DeleteQueue" awsInput:"sqs.DeleteQueueInput" awsOutput:"sqs.DeleteQueueOutput"`
	logger *logger.Logger
	graph  cloud.GraphAPI
	api    sqsiface.SQSAPI
	Url    *string `awsName:"QueueUrl" awsType:"awsstr" templateName:"url"`
}

func (cmd *DeleteQueue) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("url")))
}
