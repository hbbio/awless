package awsspec

import (
	"github.com/hbbio/awless/cloud"
	"github.com/hbbio/awless/template/params"

	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/hbbio/awless/logger"
)

type CreateRoute struct {
	_       string `action:"create" entity:"route" awsAPI:"ec2" awsCall:"CreateRoute" awsInput:"ec2.CreateRouteInput" awsOutput:"ec2.CreateRouteOutput" awsDryRun:""`
	logger  *logger.Logger
	graph   cloud.GraphAPI
	api     ec2iface.EC2API
	Table   *string `awsName:"RouteTableId" awsType:"awsstr" templateName:"table"`
	CIDR    *string `awsName:"DestinationCidrBlock" awsType:"awsstr" templateName:"cidr"`
	Gateway *string `awsName:"GatewayId" awsType:"awsstr" templateName:"gateway"`
}

func (cmd *CreateRoute) ParamsSpec() params.Spec {
	return params.NewSpec(
		params.AllOf(params.Key("cidr"), params.Key("gateway"), params.Key("table")),
		params.Validators{"cidr": params.IsCIDR})
}

type DeleteRoute struct {
	_      string `action:"delete" entity:"route" awsAPI:"ec2" awsCall:"DeleteRoute" awsInput:"ec2.DeleteRouteInput" awsOutput:"ec2.DeleteRouteOutput" awsDryRun:""`
	logger *logger.Logger
	graph  cloud.GraphAPI
	api    ec2iface.EC2API
	Table  *string `awsName:"RouteTableId" awsType:"awsstr" templateName:"table"`
	CIDR   *string `awsName:"DestinationCidrBlock" awsType:"awsstr" templateName:"cidr"`
}

func (cmd *DeleteRoute) ParamsSpec() params.Spec {
	return params.NewSpec(
		params.AllOf(params.Key("cidr"), params.Key("table")),
		params.Validators{"cidr": params.IsCIDR})
}
