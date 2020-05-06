package awsspec

import (
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/hbbio/awless/cloud"
	"github.com/hbbio/awless/logger"
	"github.com/hbbio/awless/template/params"
)

type CreateLoginprofile struct {
	_             string `action:"create" entity:"loginprofile" awsAPI:"iam" awsCall:"CreateLoginProfile" awsInput:"iam.CreateLoginProfileInput" awsOutput:"iam.CreateLoginProfileOutput"`
	logger        *logger.Logger
	graph         cloud.GraphAPI
	api           iamiface.IAMAPI
	Username      *string `awsName:"UserName" awsType:"awsstr" templateName:"username"`
	Password      *string `awsName:"Password" awsType:"awsstr" templateName:"password"`
	PasswordReset *bool   `awsName:"PasswordResetRequired" awsType:"awsbool" templateName:"password-reset"`
}

func (cmd *CreateLoginprofile) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("password"), params.Key("username"),
		params.Opt("password-reset"),
	))
}

func (cmd *CreateLoginprofile) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*iam.CreateLoginProfileOutput).LoginProfile.UserName)
}

type UpdateLoginprofile struct {
	_             string `action:"update" entity:"loginprofile" awsAPI:"iam" awsCall:"UpdateLoginProfile" awsInput:"iam.UpdateLoginProfileInput" awsOutput:"iam.UpdateLoginProfileOutput"`
	logger        *logger.Logger
	graph         cloud.GraphAPI
	api           iamiface.IAMAPI
	Username      *string `awsName:"UserName" awsType:"awsstr" templateName:"username"`
	Password      *string `awsName:"Password" awsType:"awsstr" templateName:"password"`
	PasswordReset *bool   `awsName:"PasswordResetRequired" awsType:"awsbool" templateName:"password-reset"`
}

func (cmd *UpdateLoginprofile) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("password"), params.Key("username"),
		params.Opt("password-reset"),
	))
}

type DeleteLoginprofile struct {
	_        string `action:"delete" entity:"loginprofile" awsAPI:"iam" awsCall:"DeleteLoginProfile" awsInput:"iam.DeleteLoginProfileInput" awsOutput:"iam.DeleteLoginProfileOutput"`
	logger   *logger.Logger
	graph    cloud.GraphAPI
	api      iamiface.IAMAPI
	Username *string `awsName:"UserName" awsType:"awsstr" templateName:"username"`
}

func (cmd *DeleteLoginprofile) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("username")))
}
