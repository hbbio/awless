package awsspec

import (
	"github.com/hbbio/awless/cloud"
	"github.com/hbbio/awless/template/env"
	"github.com/hbbio/awless/template/params"

	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/hbbio/awless/logger"
)

type CreateVpc struct {
	_      string `action:"create" entity:"vpc" awsAPI:"ec2" awsCall:"CreateVpc" awsInput:"ec2.CreateVpcInput" awsOutput:"ec2.CreateVpcOutput" awsDryRun:""`
	logger *logger.Logger
	graph  cloud.GraphAPI
	api    ec2iface.EC2API
	CIDR   *string `awsName:"CidrBlock" awsType:"awsstr" templateName:"cidr"`
	Name   *string `awsName:"Name" templateName:"name"`
}

func (cmd *CreateVpc) ParamsSpec() params.Spec {
	return params.NewSpec(
		params.AllOf(params.Key("cidr"), params.Opt(params.Suggested("name"))),
		params.Validators{"cidr": params.IsCIDR})
}

func (cmd *CreateVpc) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*ec2.CreateVpcOutput).Vpc.VpcId)
}

func (cmd *CreateVpc) AfterRun(renv env.Running, output interface{}) error {
	return createNameTag(awssdk.String(cmd.ExtractResult(output)), cmd.Name, renv)
}

type DeleteVpc struct {
	_      string `action:"delete" entity:"vpc" awsAPI:"ec2" awsCall:"DeleteVpc" awsInput:"ec2.DeleteVpcInput" awsOutput:"ec2.DeleteVpcOutput" awsDryRun:""`
	logger *logger.Logger
	graph  cloud.GraphAPI
	api    ec2iface.EC2API
	Id     *string `awsName:"VpcId" awsType:"awsstr" templateName:"id"`
}

func (cmd *DeleteVpc) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("id")))
}
