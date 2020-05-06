package awsspec

import (
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
	"github.com/hbbio/awless/cloud"
	"github.com/hbbio/awless/logger"
	"github.com/hbbio/awless/template/params"
)

type CreateScalingpolicy struct {
	_                   string `action:"create" entity:"scalingpolicy" awsAPI:"autoscaling" awsCall:"PutScalingPolicy" awsInput:"autoscaling.PutScalingPolicyInput" awsOutput:"autoscaling.PutScalingPolicyOutput"`
	logger              *logger.Logger
	graph               cloud.GraphAPI
	api                 autoscalingiface.AutoScalingAPI
	AdjustmentType      *string `awsName:"AdjustmentType" awsType:"awsstr" templateName:"adjustment-type"`
	Scalinggroup        *string `awsName:"AutoScalingGroupName" awsType:"awsstr" templateName:"scalinggroup"`
	Name                *string `awsName:"PolicyName" awsType:"awsstr" templateName:"name"`
	AdjustmentScaling   *int64  `awsName:"ScalingAdjustment" awsType:"awsint64" templateName:"adjustment-scaling"`
	Cooldown            *int64  `awsName:"Cooldown" awsType:"awsint64" templateName:"cooldown"`
	AdjustmentMagnitude *int64  `awsName:"MinAdjustmentMagnitude" awsType:"awsint64" templateName:"adjustment-magnitude"`
}

func (cmd *CreateScalingpolicy) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("adjustment-scaling"), params.Key("adjustment-type"), params.Key("name"), params.Key("scalinggroup"),
		params.Opt("adjustment-magnitude", "cooldown"),
	))
}

func (cmd *CreateScalingpolicy) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*autoscaling.PutScalingPolicyOutput).PolicyARN)
}

type DeleteScalingpolicy struct {
	_      string `action:"delete" entity:"scalingpolicy" awsAPI:"autoscaling" awsCall:"DeletePolicy" awsInput:"autoscaling.DeletePolicyInput" awsOutput:"autoscaling.DeletePolicyOutput"`
	logger *logger.Logger
	graph  cloud.GraphAPI
	api    autoscalingiface.AutoScalingAPI
	Id     *string `awsName:"PolicyName" awsType:"awsstr" templateName:"id"`
}

func (cmd *DeleteScalingpolicy) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("id")))
}
