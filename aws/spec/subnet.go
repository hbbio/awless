package awsspec

import (
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/hbbio/awless/cloud"
	"github.com/hbbio/awless/logger"
	"github.com/hbbio/awless/template/env"
	"github.com/hbbio/awless/template/params"
)

type CreateSubnet struct {
	_                string `action:"create" entity:"subnet" awsAPI:"ec2" awsCall:"CreateSubnet" awsInput:"ec2.CreateSubnetInput" awsOutput:"ec2.CreateSubnetOutput" awsDryRun:""`
	logger           *logger.Logger
	graph            cloud.GraphAPI
	api              ec2iface.EC2API
	CIDR             *string `awsName:"CidrBlock" awsType:"awsstr" templateName:"cidr"`
	VPC              *string `awsName:"VpcId" awsType:"awsstr" templateName:"vpc"`
	AvailabilityZone *string `awsName:"AvailabilityZone" awsType:"awsstr" templateName:"availabilityzone"`
	Public           *bool   `awsType:"awsboolattribute" templateName:"public"`
	Name             *string `templateName:"name"`
}

func (cmd *CreateSubnet) ParamsSpec() params.Spec {
	return params.NewSpec(
		params.AllOf(params.Key("cidr"), params.Key("vpc"), params.Opt(params.Suggested("name"), "availabilityzone", "public")),
		params.Validators{"cidr": params.IsCIDR})
}

func (cmd *CreateSubnet) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*ec2.CreateSubnetOutput).Subnet.SubnetId)
}

func (cmd *CreateSubnet) AfterRun(renv env.Running, output interface{}) error {
	subnetId := awssdk.String(cmd.ExtractResult(output))
	if err := createNameTag(subnetId, cmd.Name, renv); err != nil {
		return err
	}

	if BoolValue(cmd.Public) {
		updateSubnet := CommandFactory.Build("updatesubnet")().(*UpdateSubnet)
		updateSubnet.Id = subnetId
		updateSubnet.Public = Bool(true)
		if _, err := updateSubnet.Run(renv, nil); err != nil {
			return err
		}
	}

	return nil
}

type UpdateSubnet struct {
	_      string `action:"update" entity:"subnet" awsAPI:"ec2" awsCall:"ModifySubnetAttribute" awsInput:"ec2.ModifySubnetAttributeInput" awsOutput:"ec2.ModifySubnetAttributeOutput"`
	logger *logger.Logger
	graph  cloud.GraphAPI
	api    ec2iface.EC2API
	Id     *string `awsName:"SubnetId" awsType:"awsstr" templateName:"id"`
	Public *bool   `awsName:"MapPublicIpOnLaunch" awsType:"awsboolattribute" templateName:"public"`
}

func (cmd *UpdateSubnet) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("id"), params.Opt("public")))
}

type DeleteSubnet struct {
	_      string `action:"delete" entity:"subnet" awsAPI:"ec2" awsCall:"DeleteSubnet" awsInput:"ec2.DeleteSubnetInput" awsOutput:"ec2.DeleteSubnetOutput" awsDryRun:""`
	logger *logger.Logger
	graph  cloud.GraphAPI
	api    ec2iface.EC2API
	Id     *string `awsName:"SubnetId" awsType:"awsstr" templateName:"id"`
}

func (cmd *DeleteSubnet) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("id")))
}
