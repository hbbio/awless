package awsspec

import (
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/hbbio/awless/cloud"
	"github.com/hbbio/awless/logger"
	"github.com/hbbio/awless/template/params"
)

type CreateGroup struct {
	_      string `action:"create" entity:"group" awsAPI:"iam" awsCall:"CreateGroup" awsInput:"iam.CreateGroupInput" awsOutput:"iam.CreateGroupOutput"`
	logger *logger.Logger
	graph  cloud.GraphAPI
	api    iamiface.IAMAPI
	Name   *string `awsName:"GroupName" awsType:"awsstr" templateName:"name"`
}

func (cmd *CreateGroup) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("name")))
}

func (cmd *CreateGroup) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*iam.CreateGroupOutput).Group.GroupId)
}

type DeleteGroup struct {
	_      string `action:"delete" entity:"group" awsAPI:"iam" awsCall:"DeleteGroup" awsInput:"iam.DeleteGroupInput" awsOutput:"iam.DeleteGroupOutput"`
	logger *logger.Logger
	graph  cloud.GraphAPI
	api    iamiface.IAMAPI
	Name   *string `awsName:"GroupName" awsType:"awsstr" templateName:"name"`
}

func (cmd *DeleteGroup) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("name")))
}
